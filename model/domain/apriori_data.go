package domain

import (
	"apriori-backend/model/web"
	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/google/uuid"
	"strings"
)

type AprioriData struct {
	ID                 string               `gorm:"primaryKey;not null"`
	DateStart          string               `gorm:"not null"`
	DateEnd            string               `gorm:"not null"`
	MinSupport         float64              `gorm:"not null"`
	MinConfidence      float64              `gorm:"not null"`
	ItemsetSatu        []ItemsetSatu        `gorm:"foreignKey:AprioriDataID"`
	ItemsetDua         []ItemsetDua         `gorm:"foreignKey:AprioriDataID"`
	ConfidenceItemset2 []ConfidenceItemset2 `gorm:"foreignKey:AprioriDataID"`
	ItemsetTiga        []ItemsetTiga        `gorm:"foreignKey:AprioriDataID"`
	ConfidenceItemset3 []ConfidenceItemset3 `gorm:"foreignKey:AprioriDataID"`
	RuleAssociation    []RuleAssociation    `gorm:"foreignKey:AprioriDataID"`
}

func (r *AprioriData) ProceedData(apriori []Apriori.RelationRecord, request *web.CreateAprioriRequest, transactionData [][]string) *AprioriData {
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

		totalTransaction := countExactMatches(transactionData, record.GetSupportRecord().GetItems())
		switch {
		case len(record.GetSupportRecord().GetItems()) == 1:
			itemsetSatuID = uuid.New()
			r.ItemsetSatu = append(r.ItemsetSatu, ItemsetSatu{
				ID:            itemsetSatuID.String(),
				Name:          record.GetSupportRecord().GetItems()[0],
				Support:       record.GetSupportRecord().GetSupport(),
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
				confidenceItemset2ID = uuid.New()
				r.ConfidenceItemset2 = append(r.ConfidenceItemset2, ConfidenceItemset2{
					ID:            confidenceItemset2ID.String(),
					Name:          strings.Join(statistic.GetBase(), ","),
					Support:       record.GetSupportRecord().GetSupport(),
					Confidence:    statistic.GetConfidence(),
					Explanation:   "Lolos",
					AprioriDataID: r.ID,
				})

				ruleAssociationID = uuid.New()
				r.RuleAssociation = append(r.RuleAssociation, RuleAssociation{
					ID:            ruleAssociationID.String(),
					Name:          strings.Join(statistic.GetBase(), ",") + " -> " + strings.Join(statistic.GetAdd(), ","),
					Confidence:    statistic.GetConfidence(),
					LiftRatio:     statistic.GetLift(),
					Explanation:   "Lolos",
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
				r.ConfidenceItemset3 = append(r.ConfidenceItemset3, ConfidenceItemset3{
					ID:            confidenceItemset3ID.String(),
					Name:          strings.Join(statistic.GetBase(), ","),
					Support:       record.GetSupportRecord().GetSupport(),
					Confidence:    statistic.GetConfidence(),
					Explanation:   "Lolos",
					AprioriDataID: r.ID,
				})

				ruleAssociationID = uuid.New()
				r.RuleAssociation = append(r.RuleAssociation, RuleAssociation{
					ID:            ruleAssociationID.String(),
					Name:          strings.Join(statistic.GetBase(), ",") + " -> " + strings.Join(statistic.GetAdd(), ","),
					Confidence:    statistic.GetConfidence(),
					LiftRatio:     statistic.GetLift(),
					Explanation:   "Lolos",
					AprioriDataID: r.ID,
				})
			}
		}
	}
	return r
}

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
			Support:     data.Support,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ItemsetDua {
		response.ItemsetDua = append(response.ItemsetDua, web.ItemsetDua{
			Name:        strings.Split(data.Name, ","),
			Count:       data.Count,
			Support:     data.Support,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ConfidenceItemset2 {
		response.ConfidenceItemset2 = append(response.ConfidenceItemset2, web.ConfidenceItemset2{
			Name:        strings.Split(data.Name, ","),
			Support:     data.Support,
			Confidence:  data.Confidence,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ItemsetTiga {
		response.ItemsetTiga = append(response.ItemsetTiga, web.ItemsetTiga{
			Name:        strings.Split(data.Name, ","),
			Count:       data.Count,
			Support:     data.Support,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.ConfidenceItemset3 {
		response.ConfidenceItemset3 = append(response.ConfidenceItemset3, web.ConfidenceItemset3{
			Name:        strings.Split(data.Name, ","),
			Support:     data.Support,
			Confidence:  data.Confidence,
			Explanation: data.Explanation,
		})
	}

	for _, data := range r.RuleAssociation {
		response.RuleAssociation = append(response.RuleAssociation, web.RuleAssociation{
			Name:        data.Name,
			Confidence:  data.Confidence,
			LiftRatio:   data.LiftRatio,
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
