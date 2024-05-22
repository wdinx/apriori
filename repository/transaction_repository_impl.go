package repository

import (
	"apriori-backend/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (repository *TransactionRepositoryImpl) Create(transaction *domain.Transaction) error {
	if err := repository.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TransactionRepositoryImpl) Update(transaction *domain.Transaction) error {
	if err := repository.db.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TransactionRepositoryImpl) Delete(id int) error {
	if err := repository.db.Delete(&domain.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TransactionRepositoryImpl) FindById(id int) (transaction *domain.Transaction, err error) {
	if err = repository.db.First(&transaction, "id LIKE ?", id).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repository *TransactionRepositoryImpl) FindAll() (transactions *[]domain.Transaction, err error) {
	if err = repository.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repository *TransactionRepositoryImpl) FindByDateRange(startDate string, endDate string) (transactions *[]domain.Transaction, err error) {
	if err = repository.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
