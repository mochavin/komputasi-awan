# Cloud Computing Course Assignments

This repository contains assignments for the **Cloud Computing** course.

## Author
- **Name:** Moch. Avin
- **Student ID:** 5025221061

## Assignments Overview

### Assignment 1: Docker Basics
This assignment covers the fundamentals of Docker, including container management and basic operations.

📄 **Report:** [Assignment 1 - Docker Basics (PDF)](https://github.com/mochavin/komputasi-awan/blob/main/containers/docker/documentation/5025221061_Laporan%20tugas%201_%20Moch.%20Avin.pdf)

#### Cases:

| Case | Description | Technologies |
|------|-------------|--------------|
| **Case 1** | Running a background process in a container that fetches jokes from an API and saves them to a file using volume mounts | Alpine Linux, Shell Script, Volume Mounting |
| **Case 2** | Setting up a simple Python HTTP web server serving static HTML files with port publishing | Python 3, HTTP Server, Port Binding |
| **Case 3** | Running MySQL database with phpMyAdmin for database management, demonstrating container linking | MySQL 8.0, phpMyAdmin, Container Linking |
| **Case 4** | Combining a joke fetcher container with a web server to display results, demonstrating multi-container workflows | Alpine, Python, Shared Volumes |

---

### Assignment 2: Docker Compose
This assignment focuses on Docker Compose for multi-container application orchestration.

📄 **Report:** [Assignment 2 - Docker Compose (PDF)](https://github.com/mochavin/komputasi-awan/blob/main/containers/compose/compose/documentation/5025221061_Laporan%20Tugas%202_Moch.%20Avin.pdf)

#### Cases:

| Case | Description | Technologies |
|------|-------------|--------------|
| **Case 1** | Basic Nginx web server setup with custom configuration and static HTML content | Nginx, Docker Compose, Bridge Network |
| **Case 2** | Nginx web server with SSL/TLS configuration for HTTPS support using custom certificates | Nginx, SSL/TLS, HTTPS (ports 80 & 443) |
| **Case 3** | Complete WordPress stack with MySQL database, phpMyAdmin, and Nginx reverse proxy with SSL | WordPress, MySQL 8.0, phpMyAdmin, Nginx, SSL |
| **Case 4** | Full-stack PHP application with MySQL, phpMyAdmin, custom Apache build, and Alpine helper container | MySQL 5.7, Apache, PHP, phpMyAdmin, Custom Dockerfile |
| **Case 5** | Microservices architecture with PostgreSQL database, Go backend, React frontend, and Nginx API gateway | PostgreSQL, Go Backend, React Frontend, Nginx Gateway |

---

## Repository Structure

```
├── containers/
│   ├── docker/          # Assignment 1 - Docker basics
│   │   ├── case1/       # Running processes in containers
│   │   ├── case2/       # Simple web server
│   │   ├── case3/       # MySQL and phpMyAdmin
│   │   ├── case4/       # Web server with scripts
│   │   └── documentation/
│   │
│   ├── compose/         # Assignment 2 - Docker Compose
│   │   ├── compose/
│   │   │   ├── case1/   # Basic Nginx setup
│   │   │   ├── case2/   # Nginx with SSL
│   │   │   ├── case3/   # Nginx configurations
│   │   │   ├── case4/   # Full stack application
│   │   │   ├── case5/   # Microservices architecture
│   │   │   └── documentation/
│   │   └── images/
│   │
│   └── runc/
│
└── kubernetes/          # Kubernetes configurations
```

## Technologies Used

- Docker
- Docker Compose
- Nginx
- MySQL / PostgreSQL
- phpMyAdmin
- WordPress
- Python
- Go
- React
- Apache
