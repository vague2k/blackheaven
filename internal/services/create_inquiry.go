package services

import (
	"fmt"

	"github.com/vague2k/blackheaven/internal/models"
)

func CreateInquiry(inq *models.Inquiry) {
	fmt.Printf("Inquiry created!\n\n%v", inq)
}
