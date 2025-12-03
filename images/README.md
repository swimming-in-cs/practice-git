# HW4 — To-Do List Website




---

## 🧑‍💻 Features Implemented

### 1. ✅ User Authentication
- User registration  
- Login with session-based authentication  
- Guest access prevention  
- Logout support  
- Automatic creation of **default category `none`** for new users  

---

### 2. ✅ Task Management (CRUD)
Each user can:

- **Create** tasks with title, deadline, category  
- **View** tasks  
- **Update** tasks:  
  - Mark done/undone  
  - Rename task  
  - Change category  
- **Delete** tasks  

Tasks display:

- Remaining days until deadline  
- "Today" when deadline is today  
- "Overdue" if past deadline (row highlighted in red)  

---

### 3. ✅ Category Management (CRUD)
Users can:

- Create new categories  
- Rename categories  
- Delete categories (also deletes associated tasks)  

The default category:

- **Cannot be renamed**  
- **Cannot be deleted**  

---

### 4. ✅ Deadline Logic
Using SQL `DATEDIFF`:

- `days_left > 0` → *X days remaining*  
- `days_left = 0` → *Today*  
- `days_left < 0` → *Overdue* (highlighted red)  

Sorting logic:

1. Tasks **with deadlines** first  
2. Unfinished tasks before finished tasks  
3. Tasks with closer deadlines first  

---

### 5. ✅ Persistent Data
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

