package service

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"fmt"
	"strings"
	"time"

	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/go-playground/validator/v10"
)

type AprioriServiceImpl struct {
	transactionRepository    repository.TransactionRepository
	aprioriRepository        repository.AprioriRepository
	recommendationRepository repository.RecommendationRepository
	validator                *validator.Validate
}

func NewAprioriService(transactionRepository repository.TransactionRepository, aprioriRepository repository.AprioriRepository, recommendationRepository repository.RecommendationRepository, validator *validator.Validate) AprioriService {
	return &AprioriServiceImpl{
		transactionRepository:    transactionRepository,
		aprioriRepository:        aprioriRepository,
		recommendationRepository: recommendationRepository,
		validator:                validator,
	}
}

// Melakukan proses apriori berdasarkan request dari user
func (service *AprioriServiceImpl) ProcessApriori(request *web.CreateAprioriRequest) (*web.AprioriBaseResponse, error) {
	var err error
	if err = service.validator.Struct(request); err != nil {
		return nil, err
	}

	//Mengambil data transaksi berdasarkan request dari user
	result, err := service.transactionRepository.FindByDateRange(request.DateStart, request.DateEnd)
	if err != nil {
		return nil, err
	}

	if len(*result) == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	var transaction [][]string

	// Memasukkan data transaksi ke variabel transaction
	for _, column := range *result {
		// Pisahkan berdasarkan koma
		items := strings.Split(column.Items, ",")

		var cleanedItems []string
		for _, item := range items {
			// Hapus spasi jika ada
			data := strings.ReplaceAll(item, " ", "")
			cleanedItems = append(cleanedItems, data)
		}

		transaction = append(transaction, cleanedItems)
	}

	// Melakukan proses apriori terhadap data transaction
	// Set nilai min support dan min confidence berdasarkan request user
	option := Apriori.NewOptions(request.MinSup, 0.0, 0., 0.)
	apriori := Apriori.NewApriori(transaction[0:])
	aprioriResult := apriori.Calculate(option)

	// Mengolah data hasil apriori ke dalam struct yang disediakan agar lebih mudah dibaca
	var aprioriData domain.AprioriData
	proceedApriori := aprioriData.ProceedData(aprioriResult, request, transaction, request.MinConf)
	if err = service.aprioriRepository.Create(proceedApriori); err != nil {
		return nil, err
	}
	aprioriResponse := proceedApriori.ToResponse()
	return aprioriResponse, nil
}

// Mengambil semua data hasil apriori
func (service *AprioriServiceImpl) GetAll(metadata *web.Metadata) (*[]web.AprioriBaseResponse, error) {
	apriori, err := service.aprioriRepository.FindAll(metadata)

	if err != nil {
		return nil, err
	}
	var response []web.AprioriBaseResponse
	for _, data := range *apriori {
		response = append(response, *data.ToResponse())
	}
	return &response, nil
}

// Mengambil data hasil apriori berdasarkan id di database
func (service *AprioriServiceImpl) GetByID(id string) (*web.AprioriBaseResponse, error) {
	apriori, err := service.aprioriRepository.GetByID(id)
	fmt.Println("Disini")

	if err != nil {
		return nil, err
	}
	return apriori.ToResponse(), nil
}

func (service *AprioriServiceImpl) DeleteByID(id string) error {
	if err := service.aprioriRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Mengambil rekomendasi item berdasarkan hasil apriori
func (service *AprioriServiceImpl) GetRecommendationItem() (*web.RecommendationItemResponse, error) {
	recommendationRepository, err := service.recommendationRepository.GetLast()
	if err != nil {
		return nil, err
	}
	if *&recommendationRepository.Name == "" {
		return nil, fmt.Errorf("Tidak ada data transaksi dalam dua bulan terakhir")
	}
	return recommendationRepository.ToResponse(recommendationRepository.Name), nil
}

// Membuat rekomendasi item berdasarkan nilai supportnya
func (service *AprioriServiceImpl) CreateRecommendationItem() error {
	dateEnd, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	fmt.Println(dateEnd.String())
	if err != nil {
		return err
	}
	dateStart := dateEnd.AddDate(0, -2, 0)
	result, err := service.transactionRepository.FindByDateRange(dateStart.String(), dateEnd.String())
	if err != nil {
		return err
	}

	if len(*result) == 0 {
		return fmt.Errorf("Tidak ada data transaksi dalam dua bulan terakhir")
	}

	var transaction [][]string

	for _, column := range *result {
		// Pisahkan berdasarkan koma
		items := strings.Split(column.Items, ",")

		var cleanedItems []string
		for _, item := range items {
			// Hapus spasi jika ada
			data := strings.ReplaceAll(item, " ", "")
			cleanedItems = append(cleanedItems, data)
		}

		transaction = append(transaction, cleanedItems)
	}
	fmt.Println(transaction)

	option := Apriori.NewOptions(0.1, 0.1, 0., 0.)
	apriori := Apriori.NewApriori(transaction[1:])
	aprioriResult := apriori.Calculate(option)

	fmt.Println(aprioriResult)

	var recommendation domain.RecommendationItem
	var highestSupport float64
	for i := 1; i < len(aprioriResult); i++ {
		highestSupport = aprioriResult[i-1].GetSupportRecord().GetSupport()
		if aprioriResult[i].GetSupportRecord().GetSupport() >= highestSupport {
			highestSupport = aprioriResult[i].GetSupportRecord().GetSupport()
			recommendation.Name = strings.Join(aprioriResult[i].GetSupportRecord().GetItems(), ",")
			fmt.Println(highestSupport)
		}
	}
	fmt.Println("Fungsi Terpanggil")
	if err = service.recommendationRepository.Create(&recommendation); err != nil {
		return err
	}
	return nil
}

// Menghapus semua data hasil apriori
func (service *AprioriServiceImpl) DeleteAll() error {
	if err := service.aprioriRepository.DeleteAll(); err != nil {
		return err
	}
	return nil
}
