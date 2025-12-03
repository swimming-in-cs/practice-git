# HW4 — To-Do List Website




---

## Basic Requirements

### 1. registration page
- User registration  
- Login with session-based authentication  
- Guest access prevention  
- Logout support  
- Automatic creation of **default category `none`** for new users  

---

### 2.  login page
Each user can:

- **Create** tasks with title, deadline, category  
- **View** tasks  
- **Update** tasks:  
  - Mark done/undone  
  - Rename task  
  - Change category  
- **Delete** tasks  



---

### 3.  logout button in the to-do list pages
Users can:

- Create new categories  
- Rename categories  
- Delete categories (also deletes associated tasks)  


---

### 4. Add a new task, deadline and select the to-do category
Using SQL `DATEDIFF`:

- `days_left > 0` → *X days remaining*  
- `days_left = 0` → *Today*  
- `days_left < 0` → *Overdue* (highlighted red)  

Sorting logic:

1. Tasks **with deadlines** first  
2. Unfinished tasks before finished tasks  
3. Tasks with closer deadlines first  

---

### 5. Add a new category 
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 6. Change the name of the task 
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 7. Change the category of the task
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 7. Change the name of the category
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 8. Delete the task
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 9. Delete the category and the tasks of this category
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 10. Use checkbox (or other methods) to show if the task is 
finished
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 11.  Save the data to database after each modification
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 12. After reloading the page or logging in again, the content 
should not show outdated data (read data from 
database) 
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

### 8. Delete the task
- All data is stored in MySQL  
- Reloading the page always fetches **latest DB data**  
- **POST→Redirect→GET (PRG)** pattern prevents duplicate submissions  

---

## 🗄️ Database Design

### `users` table
| Field | Type | Description |
|-------|--------|-------------|
| id | INT | Primary key |
| username | VARCHAR | User login name |
| password | VARCHAR | Hashed password |

---

### `categories` table
| Field | Type | Description |
|-------|--------|-------------|
| id | INT | Primary key |
| user_id | INT | Category owner |
| name | VARCHAR | Category name |
| is_default | TINYINT(1) | 1 = default "none" |

---

### `tasks` table
| Field | Type | Description |
|-------|--------|-------------|
| id | INT | Task ID |
| user_id | INT | Task owner |
| category_id | INT | Category |
| title | VARCHAR | Task title |
| deadline | DATE / NULL | Task deadline |
| is_done | TINYINT(1) | Completion status |

---

## 🛠️ How to Install & Run

### 1. Import Database
Open **phpMyAdmin** → Import → choose `hw4.sql`.

This will create all required tables and sample data.

---

### 2. Configure `db.php`

```php
<?php
$dsn = 'mysql:host=localhost;dbname=hw4_db;charset=utf8';
$user = 'root';
$pass = '';

$pdo = new PDO($dsn, $user, $pass);
$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

session_start();
?>

