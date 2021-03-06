package repository

import (
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/models"
)

type LanguageRepository struct {}

// Create new language
func (*LanguageRepository) Create(model *models.Language) (*models.Language, error) {
    return model, database.CreateRecord(database.CreationParameters {
        Model:        model,
        CreateWithID: false,
    })
}

// Update existing language
func (*LanguageRepository) Update(model *models.Language, fields UpdatedFields) (*models.Language, error) {
    return model, database.UpdateRecord(model, fields)
}

// Soft delete existing language
func (*LanguageRepository) SoftDelete(model *models.Language) (*models.Language, error) {
    return model, database.SoftDeleteRecord(model)
}

// Force delete existing language
func (*LanguageRepository) ForceDelete(model *models.Language) (*models.Language, error) {
    return model, database.ForceDeleteRecord(model)
}

// Count total number of entries
func (*LanguageRepository) Count() (uint64, error) {
    return database.GetCount(&models.Language{})
}

// Get all languages from the database
func (*LanguageRepository) FindAll() ([]*models.Language, error) {
    var languages []*models.Language
    return languages, database.GetRecord(&database.DatabaseQuery{
        Output:     &languages,
        All:        true,
    })
}

// Find one language by ISO code
func (repository *LanguageRepository) FindOneByISO(code string) (*models.Language, error) {
    return repository.findOne(models.Language { Iso: code })
}

// Find many languages by ISO code
func (repository *LanguageRepository) FindManyByISO(code ...string) ([]*models.Language, error) {
    return repository.findMany(models.KeyLanguageIso + " IN (?)", code)
}

// Find single language
func (*LanguageRepository) findOne(query interface{}) (*models.Language, error) {
    language := &models.Language{}
    err := database.GetRecord(&database.DatabaseQuery{
        Output:     language,
        Query:      query,
    })
    return language, err
}

// Find multiple languages
func (*LanguageRepository) findMany(query interface{}, parameters interface{}) ([]*models.Language, error) {
    var languages []*models.Language
    err := database.GetRecord(&database.DatabaseQuery{
        Output:     &languages,
        Query:      query,
        Parameters: parameters,
        Multiple:   true,
    })
    return languages, err
}