package actions

import (
	"api/models"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
)

// SalesIndex default implementation.
func SalesIndex(c buffalo.Context) error {
	// offset_no := c.Params("offset_no")

	page_no := c.Param("pageNo")
	page_size := c.Param("pageSize")
	sales := []models.Order{}
	tmp1, _ := strconv.Atoi(page_no)
	tmp2, _ := strconv.Atoi(page_size)
	err := models.DB.Paginate(tmp1, tmp2).All(&sales)
	if err != nil {
		log.Panic(err)
	}
	return c.Render(http.StatusOK, r.JSON(sales))
}

// SalesStore default implementation.
func SalesStore(c buffalo.Context) error {
	salesFile, err := c.File("sales_file")
	var message int
	if err != nil {
		log.Panic(err)
	}
	// Parse the file
	reader := csv.NewReader(salesFile)
	// reader.Read()
	records, _ := reader.ReadAll()
	for _, row := range records {
		sale := &models.Order{}
		sale.Region = strings.TrimSpace(row[0])
		sale.Country = strings.TrimSpace(row[1])
		sale.ItemType = strings.TrimSpace(row[2])
		sale.SalesChannel = strings.TrimSpace(row[3])
		sale.OrderPriority = strings.TrimSpace(row[4])
		sale.OrderDate, _ = time.Parse("1/2/2006", strings.TrimSpace(row[5]))
		sale.OrderId, _ = strconv.Atoi(row[6])
		sale.ShipDate, _ = time.Parse("1/2/2006", strings.TrimSpace(row[7]))
		sale.UnitsSold, _ = strconv.Atoi(row[8])
		tmp, _ := strconv.ParseFloat(row[9], 32)
		sale.UnitPrice = float32(tmp)
		tmp, _ = strconv.ParseFloat(row[10], 32)
		sale.UnitCost = float32(tmp)
		tmp, _ = strconv.ParseFloat(row[11], 32)
		sale.TotalRevenue = float32(tmp)
		tmp, _ = strconv.ParseFloat(row[12], 32)
		sale.TotalCost = float32(tmp)
		tmp, _ = strconv.ParseFloat(row[13], 32)
		sale.TotalProfit = float32(tmp)
		err := models.DB.Create(sale)
		if err != nil {
			message = 409
			continue
		}
	}
	if message != 200 {
		if message == 409 {
			return c.Render(409, r.JSON(map[string]string{"message": "Resource Conflicts present, Possible duplicate file"}))
		}
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Succesful Operation"}))
}

//TopFiveByDate
func TopFiveByDate(c buffalo.Context) error {
	start_date := c.Param("startDate")
	end_date := c.Param("endDate")
	type result struct {
		ItemType string `db:"itemType"`
		Total    float32
	}
	itypes := []result{}
	q := models.DB.RawQuery("SELECT itemType,sum(totalProfit) as total FROM orders WHERE orderDate BETWEEN ? AND ? GROUP BY itemType ORDER BY total DESC LIMIT 5", start_date, end_date)
	err := q.All(&itypes)
	fmt.Println(err)
	return c.Render(http.StatusOK, r.JSON(itypes))
}

//ProfitInPeriod
func ProfitNPeriod(c buffalo.Context) error {
	start_date := c.Param("startDate")
	end_date := c.Param("endDate")
	type result struct {
		Total float32
	}
	totalP := result{}
	q := models.DB.RawQuery("SELECT sum(totalProfit) as total FROM orders WHERE orderDate BETWEEN ? AND ?", start_date, end_date)
	err := q.First(&totalP)
	fmt.Println(err)
	return c.Render(http.StatusOK, r.JSON(totalP))
}

func SaleCount(c buffalo.Context) error {
	type result struct {
		Total int
	}
	totalP := result{}
	q := models.DB.RawQuery("SELECT COUNT(*) as total from orders")
	err := q.First(&totalP)
	fmt.Println(err)
	return c.Render(http.StatusOK, r.JSON(totalP))
}

func SaleDates(c buffalo.Context) error {
	type result struct {
		startDate string
		endDate   string
	}
	dates := result{}
	q := models.DB.RawQuery("SELECT MIN(orderDate) AS startDate, MAX(orderDate) AS endDate FROM orders")
	err := q.First(&dates)
	if err != nil {
		fmt.Println(err)
	}
	return c.Render(http.StatusOK, r.JSON(dates))
}
