package infraestructure

import (
	"api-main/core"
	"api-main/users/domain"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

// Genera el JWT para el usuario
func generateJWT(userID int) (string, error) {
	// La clave secreta para firmar el token (debería ser más segura en producción)
	var jwtKey = []byte("eYFatÏBumBûmBac8451329*+-")

	// Definir los claims (reclamos) del JWT
	claims := jwt.MapClaims{
		"idUsuario": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	}

	// Crear el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Firmar y generar el token
	return token.SignedString(jwtKey)
}

// Save guarda un nuevo usuario con la contraseña hasheada
func (r *MySQLRepository) Save(p *domain.User) error {
	// Hashear la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO usuarios (name, user_type, last_name, email, password) VALUES (?,?,?,?,?)"
	_, err = r.conn.DB.Exec(query, &p.Name, &p.User_type, &p.Last_name, &p.Email, &hashedPassword)
	return err
}

// GetUserById obtiene un usuario por ID
func (r *MySQLRepository) GetUserById(id int) ([]domain.User, error) {
	query := "SELECT idUsuario, name, user_type, last_name, email, password FROM usuarios WHERE idUsuario = ?"
	rows, err := r.conn.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.User_type, &user.Last_name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetAllUser obtiene todos los usuarios
func (r *MySQLRepository) GetAllUser() ([]domain.User, error) {
	query := "SELECT idUsuario, name, user_type, last_name, email, password FROM usuarios"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.User_type,&user.Last_name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser actualiza un usuario
func (r *MySQLRepository) UpdateUser(id int, p *domain.User) error {
	// Hashear la nueva contraseña si se proporciona
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "UPDATE usuarios SET name=?, user_type=?,last_name=?, email=?, password=? WHERE idUsuario=?"
	_, err = r.conn.DB.Exec(query, &p.Name, &p.User_type ,&p.Last_name, &p.Email, &hashedPassword, id)
	return err
}

// DeleteUser elimina un usuario por ID
func (r *MySQLRepository) DeleteUser(id int) error {
	query := "DELETE FROM usuarios WHERE idUsuario=?"
	_, err := r.conn.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail busca un usuario por email
func (r *MySQLRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT idUsuario, name, user_type, last_name, email, password FROM usuarios WHERE email = ? LIMIT 1"
	row := r.conn.DB.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.User_type,&user.Last_name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado o error en la consulta: %w", err)
	}

	return &user, nil
}

func (r *MySQLRepository) Login(email string, password string) (*domain.User, string, error) {
	// Buscar el usuario por email
	query := "SELECT idUsuario, name, user_type, last_name, email, password FROM usuarios WHERE email = ?"
	row := r.conn.DB.QueryRow(query, email)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.User_type, &user.Last_name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, "", fmt.Errorf("usuario no encontrado")
		}
		return nil, "", fmt.Errorf("error al consultar el usuario: %w", err)
	}

	// Comparar la contraseña proporcionada con el hash almacenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", fmt.Errorf("credenciales incorrectas")
	}

	// Generar el JWT después de la validación exitosa
	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("error al generar el token: %w", err)
	}

	// Retornar el usuario y el token
	return &user, token, nil
}