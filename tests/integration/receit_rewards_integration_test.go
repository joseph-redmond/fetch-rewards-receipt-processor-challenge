package integration

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"receipt-processor-challenge/internal/receipt/model"
	"receipt-processor-challenge/internal/receipt/repository"
	"receipt-processor-challenge/internal/receipt/service"
	"receipt-processor-challenge/pkg/logger"
	"strconv"
	"strings"
	"testing"
)

type ReceiptRewardsTest struct {
	receipt      model.Receipt
	pointsEarned int64
	baseUrl      string
}

// "Given" function that will create a receipt that will score 0 points
func (t *ReceiptRewardsTest) iHaveAValidReceipt() error {
	receiptItems := []model.ReceiptItem{}
	receiptItems = append(receiptItems, model.ReceiptItem{
		"item",
		"0.01",
	})
	t.receipt = model.Receipt{
		"_",
		"2022-01-02",
		"13:33",
		"0.01",
		receiptItems,
	}
	return nil
}

// "Given" function that will set the name of the Retailer on the receipt (Can be used with other given statements)
func (t *ReceiptRewardsTest) iHaveAReceiptWithRetailerName(retailerName string) error {
	t.receipt.RetailerName = retailerName
	return nil
}

// "Given" function that will set the total on the receipt, (Can't be used without any other "Given" statements that modify the items array or the total)
func (t *ReceiptRewardsTest) iHaveAReceiptWithATotal(total string) error {
	err := t.iHaveAReceiptWithAnItem("Demo Item", total)
	return err
}

// "Given" function that will set the purchase date on the receipt (Can be used with other given statements)
func (t *ReceiptRewardsTest) iHaveAReceiptWithAPurchaseDate(date string) error {
	t.receipt.PurchaseDate = date
	return nil
}

// "Given" function that will set the purchase time on the receipt (Can be used with other given statements)
func (t *ReceiptRewardsTest) iHaveAReceiptWithAPurchaseTime(timeStr string) error {
	t.receipt.PurchaseTime = timeStr
	return nil
}

// "Given" function that will set the items array with multiple items and total (Can't be used with any other "Given" statements that set the items array or the total)
func (t *ReceiptRewardsTest) iHaveAReceiptWithItems(itemCount int, items string, total string) error {
	totalWithoutPrefix := strings.TrimPrefix(total, "$")
	totalAsFloat, err := strconv.ParseFloat(totalWithoutPrefix, 64)
	if err != nil {
		return err
	}
	splitPrice := totalAsFloat / float64(itemCount)
	priceStr := strconv.FormatFloat(splitPrice, 'f', -1, 64)

	splitItems := strings.Split(items, ",")
	itemList := []model.ReceiptItem{}
	for _, item := range splitItems {
		receiptItem := model.ReceiptItem{
			item,
			priceStr,
		}
		itemList = append(itemList, receiptItem)
	}
	t.receipt.Items = itemList
	t.receipt.TotalAmount = total
	return nil
}

// "Given" function that will set the items array with a single item and total (Can't be used with any other "Given" statements that set the items array or the total)
func (t *ReceiptRewardsTest) iHaveAReceiptWithAnItem(item string, total string) error {
	totalWithoutPrefix := strings.TrimPrefix(total, "$")

	itemList := []model.ReceiptItem{}
	newItem := model.ReceiptItem{
		item,
		totalWithoutPrefix,
	}
	itemList = append(itemList, newItem)
	t.receipt.Items = itemList
	t.receipt.TotalAmount = total
	return nil
}

// "Then" function that will check that the result from processing the receipt was a particular point value
func (t *ReceiptRewardsTest) theTotalPointsShouldBe(points string) error {
	expectedPoints, err := strconv.ParseInt(points, 10, 64)
	if err != nil {
		return err
	}
	if t.pointsEarned != expectedPoints {
		fmt.Errorf("expected %d points but got %d", expectedPoints, t.pointsEarned)
	}

	return nil
}

// "When" function that will submit the receipt for processing and save the results
func (t *ReceiptRewardsTest) iSubmitTheReceipt() error {
	theLogger := logger.GetLogger()
	receiptRepo := repository.NewRepository(theLogger)
	receiptService := service.NewService(receiptRepo, theLogger)

	processedReceipt, err := receiptService.ProcessReceipt(context.Background(), &t.receipt)

	if err != nil {
		return err
	}

	foundReceipt, err := receiptService.FindReceiptById(context.Background(), processedReceipt.ID())

	if err != nil {
		return err
	}

	t.pointsEarned = int64(foundReceipt.Points())

	return nil
}

// Initializes the testing scenario with the feature file matching statements with corresponding handlers
func InitializeScenario(ctx *godog.ScenarioContext) {
	test := &ReceiptRewardsTest{}

	ctx.Given(`I have a valid receipt with the required information`, test.iHaveAValidReceipt)
	ctx.Given(`I have a receipt with a retailer name "([^"]*)"`, test.iHaveAReceiptWithRetailerName)
	ctx.Given(`I have a receipt with a total of (\d+)`, test.iHaveAReceiptWithATotal)
	ctx.Given(`I have a receipt with (\d+) items "([^"]*)" and a final total of (\d+)`, test.iHaveAReceiptWithItems)
	ctx.Given(`I have a receipt with an item "([^"]*)" priced at (\d+)`, test.iHaveAReceiptWithAnItem)
	ctx.Given(`I have a receipt with a purchase date of "([^"]*)"`, test.iHaveAReceiptWithAPurchaseDate)
	ctx.Given(`I have a receipt with a purchase time of "([^"]*)"`, test.iHaveAReceiptWithAPurchaseTime)

	ctx.When(`I submit the receipt`, test.iSubmitTheReceipt)
	ctx.Then(`the total points should be (\d+)`, test.theTotalPointsShouldBe)
}

// Sets up the godog test suite and is the primary test that executes
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
