package repository

import "healtech-backend/server/internal/models"

func ListDocumentsRepo(userID int) ([]models.DocumentDetails, error) {
	query := `SELECT d.id, d.user_id, d.document_id, d.file_path, d.original_filename, d.file_size, d.description, d.uploaded_at, m.document_type
	          FROM documentDetails d
	          JOIN documentMaster m ON d.document_id = m.id
	          WHERE d.user_id = ?`

	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []models.DocumentDetails
	for rows.Next() {
		var doc models.DocumentDetails
		err := rows.Scan(&doc.ID, &doc.UserID, &doc.DocumentID, &doc.FilePath, &doc.OriginalFilename, &doc.FileSize, &doc.Description, &doc.UploadedAt, &doc.DocumentType)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func CreateDocumentRepo(doc *models.DocumentDetails) (int64, error) {
	query := `INSERT INTO documentDetails (user_id, document_id, file_path, original_filename, file_size, description) 
	          VALUES (?, ?, ?, ?, ?, ?)`

	result, err := DB.Exec(query, doc.UserID, doc.DocumentID, doc.FilePath, doc.OriginalFilename, doc.FileSize, doc.Description)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetDocumentByIDRepo(id int) (*models.DocumentDetails, error) {
	query := `SELECT id, user_id, document_id, file_path, original_filename, file_size, description, uploaded_at 
	          FROM documentDetails WHERE id = ?`

	row := DB.QueryRow(query, id)

	var doc models.DocumentDetails
	err := row.Scan(&doc.ID, &doc.UserID, &doc.DocumentID, &doc.FilePath, &doc.OriginalFilename, &doc.FileSize, &doc.Description, &doc.UploadedAt)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func DeleteDocumentRepo(id int) error {
	query := `DELETE FROM documentDetails WHERE id = ?`
	_, err := DB.Exec(query, id)
	return err
}
