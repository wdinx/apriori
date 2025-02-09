package domain

import (
	"apriori-backend/model/web"
	"fmt"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"

	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/google/uuid"
)

type AprioriData struct {
	gorm.Model
	ID                 string  `gorm:"primaryKey;not null"`
	DateStart          string  `gorm:"not null"`
	DateEnd            string  `gorm:"not null"`
	MinSupport         float64 `gorm:"not null"`
	MinConfidence      float64 `gorm:"not null"`
	CreatedAt          time.Time
	ItemsetSatu        []ItemsetSatu        `gorm:"foreignKey:AprioriDataID"`
	ItemsetDua         []ItemsetDua         `gorm:"foreignKey:AprioriDataID"`
	ConfidenceItemset2 []ConfidenceItemset2 `gorm:"foreignKey:AprioriDataID"`
	ItemsetTiga        []ItemsetTiga        `gorm:"foreignKey:AprioriDataID"`
	ConfidenceItemset3 []ConfidenceItemset3 `gorm:"foreignKey:AprioriDataID"`
	RuleAssociation    []RuleAssociation    `gorm:"foreignKey:AprioriDataID"`
}

// Memproses data transaksi dan menyimpannya ke struct AprioriData
func (r *AprioriData) ProceedData(apriori []Apriori.RelationRecord, request *web.CreateAprioriRequest, transactionData [][]string, minConf float64) *AprioriData {
	r.ID = uuid.New().String()
	r.DateStart = request.DateStart
	r.DateEnd = request.DateEnd
	r.MinSupport = request.MinSup
	r.MinConfidence = request.MinConf
	var itemsetSatuID, itemsetDuaID, itemsetTigaID, ruleAssociationID, confidenceItemset2ID, confidenceItemset3ID uuid.UUID
	for _, record := range apriori {
		if len(record.GetSupportRecord().GetItems()) > 3 {
			continue
		}

		fmt.Print(record.GetSupportRecord().GetSupport())

		totalTransaction := countExactMatches(transactionData, record.GetSupportRecord().GetItems())
		switch {
		case len(record.GetSupportRecord().GetItems()) == 1:
			itemsetSatuID = uuid.New()
			r.ItemsetSatu = append(r.ItemsetSatu, ItemsetSatu{
				ID:            itemsetSatuID.String(),
				Name:          record.GetSupportRecord().GetItems()[0],
				Support:       math.Round(record.GetSupportRecord().GetSupport()*100) / 100,
				Count:         totalTransaction,
				Explanation:   "Lolos",
				AprioriDataID: r.ID,
			})
		case len(record.GetSupportRecord().GetItems()) == 2:
			itemsetDuaID = uuid.New()

			r.ItemsetDua = append(r.ItemsetDua, ItemsetDua{
				ID:            itemsetDuaID.String(),
				Name:          strings.Join(record.GetSupportRecord().GetItems(), ","),
				Support:       record.GetSupportRecord().GetSupport(),
				Count:         totalTransaction,
				Explanation:   "Lolos",
				AprioriDataID: r.ID,
			})

			for _, statistic := range record.GetOrderedStatistic() {
				// Apabila ingin menampilkan confidence itemset 2

				confidenceItemset2ID = uuid.New()
				var confidenceExplanation string
				if minConf <= statistic.GetConfidence() {
					confidenceExplanation = "Lolos"
				} else {
					confidenceExplanation = "Tidak Lolos"
				}

				r.ConfidenceItemset2 = append(r.ConfidenceItemset2, ConfidenceItemset2{
					ID:            confidenceItemset2ID.String(),
					Name:          strings.Join(statistic.GetBase(), ","),
					Support:       record.GetSupportRecord().GetSupport(),
					Confidence:    statistic.GetConfidence(),
					Explanation:   confidenceExplanation,
					AprioriDataID: r.ID,
				})

				if statistic.GetConfidence() < minConf {
					continue
				}

				ruleAssociationID = uuid.New()
				var explanation string
				if statistic.GetLift() >= 1 {
					explanation = "Positif"
				} else {
					explanation = "Negatif"
				}
				r.RuleAssociation = append(r.RuleAssociation, RuleAssociation{
					ID:            ruleAssociationID.String(),
					Name:          strings.Join(statistic.GetBase(), ",") + " -> " + strings.Join(statistic.GetAdd(), ","),
					Confidence:    statistic.GetConfidence(),
					LiftRatio:     statistic.GetLift(),
					Explanation:   explanation,
					AprioriDataID: r.ID,
				})
			}
		case len(record.GetSupportRecord().GetItems()) == 3:
			itemsetTigaID = uuid.New()
			r.ItemsetTiga = append(r.ItemsetTiga, ItemsetTiga{
				ID:            itemsetTigaID.String(),
				Name:          strings.Join(record.GetSupportRecord().GetItems(), ","),
				Support:       record.GetSupportRecord().GetSupport(),
				Count:         totalTransaction,
				Explanation:   "Lolos",
				AprioriDataID: r.ID,
			})
			for _, statistic := range record.GetOrderedStatistic() {
				confidenceItemset3ID = uuid.New()

				var confidenceExplanation string
				if minConf <= statistic.GetConfidence() {
					confidenceExplanation = "Lolos"
				} else {
					confidenceExplanation = "Tidak Lolos"
				}
				r.ConfidenceItemset3 = append(r.ConfidenceItemset3, ConfidenceItemset3{
					ID:            confidenceItemset3ID.String(),
					Name:          strings.Join(statistic.GetBase(), ","),
					Support:       record.GetSupportRecord().GetSupport(),
					Confidence:    statistic.GetConfidence(),
					Explanation:   confidenceExplanation,
					AprioriDataID: r.ID,
				})

				if statistic.GetConfidence() < minConf {
					continue
				}

				ruleAssociationID = uuid.New()
				var explanation string
				if statistic.GetLift() >= 1 {
					explanation = "Positif"
				} else {
					explanation = "Negatif"
				}
				r.RuleAssociation = append(r.RuleAssociation, RuleAssociation{
					ID:            ruleAssociationID.String(),
					Name:          strings.Join(statistic.GetBase(), ",") + " -> " + strings.Join(statistic.GetAdd(), ","),
					Confidence:    statistic.GetConfidence(),
					LiftRatio:     statistic.GetLift(),
					Explanation:   explanation,
					AprioriDataID: r.ID,
				})
			}
		}
	}
	return r
}

// Memproses data apriori agar lebih mudah dibaca dan mengirimnya ke controller untuk ditampilkan ke User.
func (r *AprioriData) ToResponse() *web.AprioriBaseResponse {
	var response web.AprioriBaseResponse

	response.ID = r.ID
	response.DateStart = r.DateStart
	response.DateEnd = r.DateEnd
	response.MinSupport = r.MinSupport
	response.MinConfidence = r.MinConfidence

	for _, data := range r.ItemsetSatu {
		response.ItemsetSatu = append(response.ItemsetSatu, web.ItemsetSatu{
			Name:        strings.Split(data.Name, ","),
			Count:       data.Count,
			Support:     math.Round(data.Support*100) / 100,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ItemsetDua {
		response.ItemsetDua = append(response.ItemsetDua, web.ItemsetDua{
			Name:        strings.Split(data.Name, ","),
			Count:       data.Count,
			Support:     math.Round(data.Support*100) / 100,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ConfidenceItemset2 {
		response.ConfidenceItemset2 = append(response.ConfidenceItemset2, web.ConfidenceItemset2{
			Name:        strings.Split(data.Name, ","),
			Support:     math.Round(data.Support*100) / 100,
			Confidence:  math.Round(data.Confidence*100) / 100,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ItemsetTiga {
		response.ItemsetTiga = append(response.ItemsetTiga, web.ItemsetTiga{
			Name:        strings.Split(data.Name, ","),
			Count:       data.Count,
			Support:     math.Round(data.Support*100) / 100,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ConfidenceItemset3 {
		response.ConfidenceItemset3 = append(response.ConfidenceItemset3, web.ConfidenceItemset3{
			Name:        strings.Split(data.Name, ","),
			Support:     math.Round(data.Support*100) / 100,
			Confidence:  math.Round(data.Confidence*100) / 100,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.RuleAssociation {
		response.RuleAssociation = append(response.RuleAssociation, web.RuleAssociation{
			Name:        data.Name,
			Confidence:  math.Round(data.Confidence*100) / 100,
			LiftRatio:   math.Round(data.LiftRatio*100) / 100,
			Explanation: data.Explanation,
		})
	}

	return &response
}

func countExactMatches(data [][]string, query []string) int {
	count := 0

	// Iterasi melalui setiap elemen dalam `data`
	for _, group := range data {
		matches := 0

		// Periksa apakah setiap elemen dalam `query` ada di `group`
		for _, q := range query {
			for _, val := range group {
				if q == val {
					matches++
					break // Lanjut ke elemen query berikutnya
				}
			}
		}

		// Jika semua elemen dalam `query` ada di `group`, tambahkan count
		if matches == len(query) {
			count++
		}
	}

	return count
}
