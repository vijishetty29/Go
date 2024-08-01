package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vijishetty29/Go/go-sitcon/pkg/dto"
	"github.com/vijishetty29/Go/go-sitcon/pkg/models"
)

const timestampLayout = "2006-01-02T15:04:05.000-0700"

func GetSelectionsData(c *gin.Context) {

	lastFetchTimestamp := c.Query("lastFetchTimestamp")
	fullLoad := c.Query("fullLoad")
	fullLoadBool, err := strconv.ParseBool(fullLoad)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "fullLoad type not supported"})
		return
	}

	if fullLoadBool {
		c.JSON(http.StatusOK, gin.H{"requestId": uuid.New(), "requestStatus": "PROCESSING"})
		return
	}

	since, parseError := time.Parse(timestampLayout, lastFetchTimestamp)
	fmt.Println("lastFetchTimestamp=", since)
	fmt.Println("parseError=", parseError)

	activeSelectionInts, _ := models.GetActiveSelections()
	fmt.Println("activeSelectionInts=", activeSelectionInts)

	selectionEntities, _ := models.GetModifiedSelections(since, activeSelectionInts)
	selectionTopicEntities, _ := models.GetModifiedSelectionTopics(since, activeSelectionInts)
	selectionItemEntities, _ := models.GetModifiedSelectionItems(since)

	// Create channels of proper size to hold conversion result
	// Create smaller chunks of SelectionItems
	// Create Channels for smaller chunks of SelectionItems
	// Create product channels, each channel holds conversion result of maximum 500 products
	demostrateGoRoutine, _ := strconv.ParseBool(os.Getenv("DEMO_GOROUTINE"))
	if demostrateGoRoutine {
		selectionData := processRequestUsingGoRoutines(selectionEntities, selectionTopicEntities, selectionItemEntities)
		c.JSON(http.StatusOK, selectionData)
	} else {
		selectionData := processRequestWithoutGoRoutines(selectionEntities, selectionTopicEntities, selectionItemEntities)
		c.JSON(http.StatusOK, selectionData)
	}
}

func processRequestUsingGoRoutines(selectionEntities []*models.SelectionEntity, selectionTopicEntities []*models.SelectionTopicEntity, selectionItemEntities []*models.SelectionItemEntity) dto.SelectionDataDTO {
	channelCounter := 0

	var selectionsChannel chan []dto.SelectionDTO
	if len(selectionEntities) != 0 {
		selectionsChannel = make(chan []dto.SelectionDTO, len(selectionEntities))
		channelCounter++
	}

	var selectionTopicsChannel chan []dto.SelectionTopicDTO
	if len(selectionTopicEntities) != 0 {
		selectionTopicsChannel = make(chan []dto.SelectionTopicDTO, len(selectionTopicEntities))
		channelCounter++
	}

	batchSize := 500
	numOfBatches := len(selectionItemEntities) / batchSize
	smallerSelectionItemsSlices := make([][]*models.SelectionItemEntity, numOfBatches)
	for i := 0; i < numOfBatches; i++ {
		smallerSlice := selectionItemEntities[i*batchSize : (i+1)*batchSize]
		smallerSelectionItemsSlices[i] = smallerSlice
	}

	var productsChannels []chan []dto.SelectionItemDTO
	for _, smallerSlice := range smallerSelectionItemsSlices {
		productsChannel := make(chan []dto.SelectionItemDTO, len(smallerSlice))
		productsChannels = append(productsChannels, productsChannel)
		channelCounter++
	}

	println("channelCounter=", channelCounter)

	var waitGroup sync.WaitGroup
	waitGroup.Add(channelCounter)

	for i := 0; i < numOfBatches; i++ {
		go GetModifiedProducts(smallerSelectionItemsSlices[i], productsChannels[i], &waitGroup)
	}

	go GetModifiedTopics(selectionTopicEntities, selectionTopicsChannel, &waitGroup)
	go GetModifiedSelections(selectionEntities, selectionsChannel, &waitGroup)

	println("All go routines triggered")
	waitGroup.Wait()

	selections := <-selectionsChannel
	topics := <-selectionTopicsChannel

	var allProducts []dto.SelectionItemDTO
	for i := 0; i < numOfBatches; i++ {
		subProducts := <-productsChannels[i]
		allProducts = append(allProducts, subProducts...)
	}

	fmt.Println("selectionSize=", len(selections))

	selectionData := dto.SelectionDataDTO{
		RequestId:     uuid.New().String(),
		RequestStatus: "COMPLETED",
		Selections:    selections,
		Topics:        topics,
		Products:      allProducts,
	}
	return selectionData
}

func processRequestWithoutGoRoutines(selectionEntities []*models.SelectionEntity, selectionTopicEntities []*models.SelectionTopicEntity, selectionItemEntities []*models.SelectionItemEntity) dto.SelectionDataDTO {

	//fmt.Println("selectionSize=", len(selections))

	selections := GetModifiedSelectionsWithoutRoutine(selectionEntities)
	topics := GetModifiedTopicsWithoutGoRoutines(selectionTopicEntities)
	products := GetModifiedProductsWithoutGoRoutines(selectionItemEntities)

	selectionData := dto.SelectionDataDTO{
		RequestId:     uuid.New().String(),
		RequestStatus: "COMPLETED",
		Selections:    selections,
		Topics:        topics,
		Products:      products,
	}
	return selectionData
}

func GetModifiedSelections(selectionEntities []*models.SelectionEntity, resultChannel chan<- []dto.SelectionDTO, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	selections := []dto.SelectionDTO{}
	for _, selection := range selectionEntities {

		active := false
		if selection.Active == 1 {
			active = true
		}

		deliveryEndWeek := dto.YearAndWeek{}
		if selection.DeliveryEndWeek != 0 && selection.DeliveryEndYear != 0 {
			deliveryEndWeek = dto.YearAndWeek{
				Week: selection.DeliveryEndWeek,
				Year: selection.DeliveryEndYear,
			}
		}

		deliveryStartWeek := dto.YearAndWeek{}
		if selection.DeliveryStartWeek != 0 && selection.DeliveryStartYear != 0 {
			deliveryStartWeek = dto.YearAndWeek{
				Week: selection.DeliveryStartWeek,
				Year: selection.DeliveryStartYear,
			}
		}

		advertisingStartWeek := dto.YearAndWeek{}
		if selection.AdvertisingStartWeek != 0 && selection.AdvertisingStartYear != 0 {
			advertisingStartWeek = dto.YearAndWeek{
				Week: selection.AdvertisingStartWeek,
				Year: selection.AdvertisingStartYear,
			}
		}

		advertisingEndWeek := dto.YearAndWeek{}
		if selection.AdvertisingEndWeek != 0 && selection.AdvertisingEndYear != 0 {
			advertisingEndWeek = dto.YearAndWeek{
				Week: selection.AdvertisingEndWeek,
				Year: selection.AdvertisingEndYear,
			}
		}

		sel := dto.SelectionDTO{
			Code:                 selection.Code,
			Active:               active,
			Branch:               "NON_FOOD",
			AdvertisingStartWeek: &advertisingStartWeek,
			AdvertisingEndWeek:   &advertisingEndWeek,
			DeliveryStartWeek:    &deliveryStartWeek,
			DeliveryEndWeek:      &deliveryEndWeek,
			// Assets:nil
		}
		selections = append(selections, sel)
	}
	resultChannel <- selections
	close(resultChannel)
}

func GetModifiedSelectionsWithoutRoutine(selectionEntities []*models.SelectionEntity) []dto.SelectionDTO {

	selections := []dto.SelectionDTO{}
	for _, selection := range selectionEntities {

		active := false
		if selection.Active == 1 {
			active = true
		}

		deliveryEndWeek := dto.YearAndWeek{}
		if selection.DeliveryEndWeek != 0 && selection.DeliveryEndYear != 0 {
			deliveryEndWeek = dto.YearAndWeek{
				Week: selection.DeliveryEndWeek,
				Year: selection.DeliveryEndYear,
			}
		}

		deliveryStartWeek := dto.YearAndWeek{}
		if selection.DeliveryStartWeek != 0 && selection.DeliveryStartYear != 0 {
			deliveryStartWeek = dto.YearAndWeek{
				Week: selection.DeliveryStartWeek,
				Year: selection.DeliveryStartYear,
			}
		}

		advertisingStartWeek := dto.YearAndWeek{}
		if selection.AdvertisingStartWeek != 0 && selection.AdvertisingStartYear != 0 {
			advertisingStartWeek = dto.YearAndWeek{
				Week: selection.AdvertisingStartWeek,
				Year: selection.AdvertisingStartYear,
			}
		}

		advertisingEndWeek := dto.YearAndWeek{}
		if selection.AdvertisingEndWeek != 0 && selection.AdvertisingEndYear != 0 {
			advertisingEndWeek = dto.YearAndWeek{
				Week: selection.AdvertisingEndWeek,
				Year: selection.AdvertisingEndYear,
			}
		}

		sel := dto.SelectionDTO{
			Code:                 selection.Code,
			Active:               active,
			Branch:               "NON_FOOD",
			AdvertisingStartWeek: &advertisingStartWeek,
			AdvertisingEndWeek:   &advertisingEndWeek,
			DeliveryStartWeek:    &deliveryStartWeek,
			DeliveryEndWeek:      &deliveryEndWeek,
			// Assets:nil
		}
		selections = append(selections, sel)
	}
	return selections

}

func GetModifiedTopics(selectionTopicEntities []*models.SelectionTopicEntity, resultChannel chan<- []dto.SelectionTopicDTO, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	topics := []dto.SelectionTopicDTO{}
	for _, topic := range selectionTopicEntities {
		statusBool := GetStatusAsBoolean(topic)

		description := models.GetDescriptionForSelectionTopicAndLanguage(topic.ID, "8796093055008")

		desc := dto.LocalizedStringDTO{
			DE: description,
		}

		calendarWeeks := models.GetCalendarWeeksForSelectionTopic(topic.ID)

		calendarWeeksData := []dto.YearAndWeek{}

		for _, calendarWeek := range calendarWeeks {
			year, _ := strconv.Atoi(calendarWeek.Year)
			week, _ := strconv.Atoi(calendarWeek.Week)
			calendarWeekData := dto.YearAndWeek{
				Year: year,
				Week: week,
				ID:   calendarWeek.ID,
			}
			calendarWeeksData = append(calendarWeeksData, calendarWeekData)
		}

		topic := dto.SelectionTopicDTO{
			Code:             topic.Code,
			Status:           statusBool,
			TopicNr:          topic.TopicNr,
			Branch:           "NON_FOOD",
			Selection:        topic.Selection.Code,
			Description:      &desc,
			AdvertisingWeeks: calendarWeeksData,
			// Assets:nil
		}
		topics = append(topics, topic)
	}
	resultChannel <- topics
	close(resultChannel)
}

func GetModifiedTopicsWithoutGoRoutines(selectionTopicEntities []*models.SelectionTopicEntity) []dto.SelectionTopicDTO {
	topics := []dto.SelectionTopicDTO{}
	for _, topic := range selectionTopicEntities {
		statusBool := GetStatusAsBoolean(topic)

		description := models.GetDescriptionForSelectionTopicAndLanguage(topic.ID, "8796093055008")

		desc := dto.LocalizedStringDTO{
			DE: description,
		}

		calendarWeeks := models.GetCalendarWeeksForSelectionTopic(topic.ID)

		calendarWeeksData := []dto.YearAndWeek{}

		for _, calendarWeek := range calendarWeeks {
			year, _ := strconv.Atoi(calendarWeek.Year)
			week, _ := strconv.Atoi(calendarWeek.Week)
			calendarWeekData := dto.YearAndWeek{
				Year: year,
				Week: week,
				ID:   calendarWeek.ID,
			}
			calendarWeeksData = append(calendarWeeksData, calendarWeekData)
		}

		topic := dto.SelectionTopicDTO{
			Code:             topic.Code,
			Status:           statusBool,
			TopicNr:          topic.TopicNr,
			Branch:           "NON_FOOD",
			Selection:        topic.Selection.Code,
			Description:      &desc,
			AdvertisingWeeks: calendarWeeksData,
			// Assets:nil
		}
		topics = append(topics, topic)
	}
	return topics
}

func GetStatusAsBoolean(topic *models.SelectionTopicEntity) bool {
	status, err := models.GetEnumerationByID(topic.Status)
	statusBool := false
	if err != nil {
		statusBool = false
	}
	if status.CodeLowerCase == "active" {
		statusBool = true
	}
	return statusBool
}

func GetSelectionCodeForSelection(selectionModels []*models.SelectionEntity, pk int64) string {
	for _, selection := range selectionModels {
		if selection.PK == pk {
			return selection.Code
		}
	}
	return ""
}

func GetModifiedProductsWithoutGoRoutines(selectionItemEntities []*models.SelectionItemEntity) []dto.SelectionItemDTO {

	products := []dto.SelectionItemDTO{}
	for _, item := range selectionItemEntities {
		code := getProductCode(item)

		localizedProdData := models.GetProductLocalizedEntityForProductAndLanguage(item.ProductPK, "8796093055008")
		var localizedName dto.LocalizedStringDTO
		var localizedMktText string
		if localizedProdData != nil {
			localizedName = dto.LocalizedStringDTO{
				DE: localizedProdData.Name,
			}
			localizedMktText = localizedProdData.RawMarketingText
		}

		psaStatus := models.GetPSAStatusForWorkflowDocumentAndCountryOrderStatus(item.WorkflowStatus, item.IdocStatus, item.CountryOrderStatus)
		var itemStatus dto.ItemStatusDTO
		if psaStatus != 0 {
			itemStatus = dto.ItemStatusDTO{
				Name: "psaStatus",
			}
		}

		assemblyType, _ := models.GetPimEnumerationCodeByID(item.Product.AssemblyType)
		assembly := dto.AssemblyDTO{
			ContainerSize: item.Product.PimContainerSize,
			ItemCount:     item.Product.AssemblyItemCount,
			Type:          assemblyType,
		}

		productCategoryCode := ""
		productCategory, _ := models.GetEnumerationByID(item.ProductCategory)
		if productCategory != nil {
			productCategoryCode = productCategory.Code
		}

		buyerName, _ := models.GetBuyerNameByID(item.Buyer)
		var buyerDTO dto.BuyerDTO
		if buyerName != "" {
			buyerDTO = dto.BuyerDTO{
				FullName: buyerName,
			}
		}

		articleClassification, _ := models.GetPimEnumerationCodeByID(item.ArticleClassification)
		characteristics := dto.CharacteristicsDTO{
			RepeatingArticle:  getBooleanValue(item.Repeating),
			DisplayArticle:    getBooleanValue(item.Product.IsDisplayItem),
			InnovationArticle: getBooleanValue(item.InnovationArticle),
			BriefingHL:        getBooleanValue(item.BriefingHl),
			CatalogArticle:    getBooleanValue(item.CatalogArticle),
			Classification:    articleClassification,
			CountryProposal:   item.CountryProposal,
		}

		var displayType dto.DisplayTypeDTO
		displayTypeCode, _ := models.GetPimEnumerationCodeByID(item.Product.CbxDisplayType)
		if displayTypeCode != "" {
			localizedDisplayType := dto.LocalizedStringDTO{
				DE: displayTypeCode,
			}
			displayType = dto.DisplayTypeDTO{
				Name: localizedDisplayType,
			}
		}

		articleType, _ := models.GetPimEnumerationCodeByID(item.ArticleType)

		product := dto.SelectionItemDTO{
			Code:             code,
			SelectionItemID:  item.Code,
			Name:             localizedName,
			Topic:            item.TopicCode,
			Status:           itemStatus,
			PurchasingRemark: item.Note,
			//OrderableCountries:       nil,
			// DeliveryDates: nil,
			//Brand:  nil,
			Branch:          "NON_FOOD",
			ProductCategory: productCategoryCode,
			Buyer:           buyerDTO,
			Characteristics: characteristics,
			DisplayType:     displayType,
			// ItemGroup: nil,
			Assembly: assembly,
			ItemType: articleType,
			//Predecessor:nil
			AdditionalDescriptionInt: localizedMktText,
			// LinkedItem:nil,
			// Asset Data,
		}
		products = append(products, product)
	}
	return products
}

func GetModifiedProducts(selectionItemEntities []*models.SelectionItemEntity, resultChannel chan<- []dto.SelectionItemDTO, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	products := []dto.SelectionItemDTO{}
	for _, item := range selectionItemEntities {
		code := getProductCode(item)

		localizedProdData := models.GetProductLocalizedEntityForProductAndLanguage(item.ProductPK, "8796093055008")
		var localizedName dto.LocalizedStringDTO
		var localizedMktText string
		if localizedProdData != nil {
			localizedName = dto.LocalizedStringDTO{
				DE: localizedProdData.Name,
			}
			localizedMktText = localizedProdData.RawMarketingText
		}

		psaStatus := models.GetPSAStatusForWorkflowDocumentAndCountryOrderStatus(item.WorkflowStatus, item.IdocStatus, item.CountryOrderStatus)
		var itemStatus dto.ItemStatusDTO
		if psaStatus != 0 {
			itemStatus = dto.ItemStatusDTO{
				Name: "psaStatus",
			}
		}

		assemblyType, _ := models.GetPimEnumerationCodeByID(item.Product.AssemblyType)
		assembly := dto.AssemblyDTO{
			ContainerSize: item.Product.PimContainerSize,
			ItemCount:     item.Product.AssemblyItemCount,
			Type:          assemblyType,
		}

		productCategoryCode := ""
		productCategory, _ := models.GetEnumerationByID(item.ProductCategory)
		if productCategory != nil {
			productCategoryCode = productCategory.Code
		}

		buyerName, _ := models.GetBuyerNameByID(item.Buyer)
		var buyerDTO dto.BuyerDTO
		if buyerName != "" {
			buyerDTO = dto.BuyerDTO{
				FullName: buyerName,
			}
		}

		articleClassification, _ := models.GetPimEnumerationCodeByID(item.ArticleClassification)
		characteristics := dto.CharacteristicsDTO{
			RepeatingArticle:  getBooleanValue(item.Repeating),
			DisplayArticle:    getBooleanValue(item.Product.IsDisplayItem),
			InnovationArticle: getBooleanValue(item.InnovationArticle),
			BriefingHL:        getBooleanValue(item.BriefingHl),
			CatalogArticle:    getBooleanValue(item.CatalogArticle),
			Classification:    articleClassification,
			CountryProposal:   item.CountryProposal,
		}

		var displayType dto.DisplayTypeDTO
		displayTypeCode, _ := models.GetPimEnumerationCodeByID(item.Product.CbxDisplayType)
		if displayTypeCode != "" {
			localizedDisplayType := dto.LocalizedStringDTO{
				DE: displayTypeCode,
			}
			displayType = dto.DisplayTypeDTO{
				Name: localizedDisplayType,
			}
		}

		articleType, _ := models.GetPimEnumerationCodeByID(item.ArticleType)

		product := dto.SelectionItemDTO{
			Code:             code,
			SelectionItemID:  item.Code,
			Name:             localizedName,
			Topic:            item.TopicCode,
			Status:           itemStatus,
			PurchasingRemark: item.Note,
			//OrderableCountries:       nil,
			// DeliveryDates: nil,
			//Brand:  nil,
			Branch:          "NON_FOOD",
			ProductCategory: productCategoryCode,
			Buyer:           buyerDTO,
			Characteristics: characteristics,
			DisplayType:     displayType,
			// ItemGroup: nil,
			Assembly: assembly,
			ItemType: articleType,
			//Predecessor:nil
			AdditionalDescriptionInt: localizedMktText,
			// LinkedItem:nil,
			// Asset Data,
		}
		products = append(products, product)
	}
	resultChannel <- products
	close(resultChannel)
}

func getBooleanValue(val int) bool {
	return val == 1
}

func getProductCode(item *models.SelectionItemEntity) string {
	if item.Product.CbxUniqueStyleId != "" {
		return item.Product.CbxUniqueStyleId
	} else if item.Product.ItemNumber != "" {
		return item.Product.ItemNumber
	} else {
		return item.Product.Code
	}
}
