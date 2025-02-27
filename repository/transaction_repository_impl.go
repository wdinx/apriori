package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

// Menyimpan data transaksi ke dalam database
func (repository *TransactionRepositoryImpl) Create(transaction *domain.Transaction) error {
	if err := repository.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

// Update data transaksi di database
func (repository *TransactionRepositoryImpl) Update(transaction *domain.Transaction) error {
	if err := repository.db.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}

// Menghapus data transaksi di database berdasarkan ID nya
func (repository *TransactionRepositoryImpl) Delete(id int) error {
	if err := repository.db.Delete(&domain.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Mengambil data transaksi berdasarkan ID nya
func (repository *TransactionRepositoryImpl) FindById(id int) (transaction *domain.Transaction, err error) {
	if err = repository.db.First(&transaction, "id LIKE ?", id).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// Mengambil data transaksi berdasakan request User
func (repository *TransactionRepositoryImpl) FindAll(metadata *web.Metadata) (transactions *[]domain.Transaction, err error) {
	if err = repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Mengambil data transaksi berdasarkan range tanggal yang request user
func (repository *TransactionRepositoryImpl) FindByDateRange(startDate string, endDate string) (transactions *[]domain.Transaction, err error) {
	if err = repository.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Mengambil semua data transaksi di database
func (repository *TransactionRepositoryImpl) GetALl() (*[]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := repository.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return &transactions, nil
}

// Menyimpan data transaksi ke dalam database berdasarkan file excel yang di input oleh User
func (repository *TransactionRepositoryImpl) InsertByExcel(transaction *[]domain.Transaction) error {
	if err := repository.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TransactionRepositoryImpl) DeleteAll() error {
	if err := repository.db.Exec("DELETE FROM transactions").Error; err != nil {
		return err
	}
	return nil
}
