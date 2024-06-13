package handlers

import (
	"basicGoCrud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create item
func CreateItem(db *gorm.DB) gin.HandlerFunc { //gorm.db sử dụng con trỏ để làm việc với cùng một kết nối cơ sở dữ liệu trong suốt quá trình thực thi.
	return func(c *gin.Context) { //gin.Context sử dụng con trỏ để thao tác trên cùng một đối tượng Context cho tất cả các hàm
		var newItem models.TodoItem

		//Parse dữ liệu json vào biến newItem
		if err := c.ShouldBindJSON(&newItem); err != nil { //Sử dụng &newItem cho phép hàm này cập nhật trực tiếp vào newItem gốc, thay vì làm việc với một bản sao. (tham chiếu)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Check if item title is already exist then return message "item title already exist"
		var existingItem models.TodoItem
		result := db.Where("title = ?", newItem.Title).First(&existingItem)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Item title already exist"})
			return
		}

		if result := db.Create(&newItem); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, "Created successfully!!")
	}
}

// Get item by id
func GetItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.TodoItem
		if result := db.First(&item, c.Param("id")); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, item)
	}
}

// Get all item
func GetItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var items []models.TodoItem
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))    //Lấy tham số page nếu ko có sẽ là 1
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20")) //Lấy tham số limit nếu ko có sẽ là 20
		offset := (page - 1) * limit                            //Số lượng bảng ghi mà truy vấn sẽ bỏ qua. Ví dụ: Trang 2 limit 5 thì nó sẽ truy vấn (2-1) * 5 nghĩa là 5 item tiếp theo bỏ qua trang 1

		if result := db.Limit(limit).Offset(offset).Find(&items); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, items)
	}
}

func UpdateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.TodoItem
		if result := db.First(&item, c.Param("id")); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Save(&item); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, item)
	}
}

func DeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if result := db.Delete(&models.TodoItem{}, c.Param("id")); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	}
}
