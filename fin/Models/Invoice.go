package models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	BaseSoftDel
	Technology_Product_ID uuid.UUID
	Invoice_Date          time.Time
	total_price           uint
}
