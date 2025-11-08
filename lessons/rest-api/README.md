# Project: REST API (Go)

A simple REST API built with Go — focusing on clean structure and clear learning steps.

---

## 📌 Project Overview

This project demonstrates how to build a REST API using the Go programming language.  
It covers endpoint handling, request/response processing, authentication, and using JSON for data exchange.

---

## 🗂️ Project Steps

1. **Planning** the API design  
2. **Building** the REST API step-by-step  
3. Creating endpoints  
4. Handling HTTP methods  
5. Adding authentication (JWT)  
6. Implementing access control  

---

## 🤔 What is a REST API?

A **REST API** allows a **Client** (e.g., laptop, browser, phone) to communicate with a **Server**.

| Client | → Sends Request | Server | → Sends Response |
|-------|----------------|--------|------------------|
| Laptop / Phone | GET `/some-data` | Go Backend | JSON Data |
| (or) | POST `/some-data` | | Success / Error |

### Key Concepts:
- Data is exchanged in **JSON** format.
- Each operation is defined by an **HTTP Method + URL Endpoint**.

---

## 🎫 Project Description: Event Booking REST API

A Go-powered **Event Booking** REST API.

| Method | Endpoint | Description | Auth |
|-------|----------|-------------|------|
| GET | `/events` | Get list of available events | ❌ Not Required |
| GET | `/events/{id}` | Get details of a specific event | ❌ Not Required |
| POST | `/events` | Create a new bookable event | ✅ Required |
| PUT | `/events/{id}` | Update an event (only creator) | ✅ Required |
| DELETE | `/events/{id}` | Delete an event (only creator) | ✅ Required |
| POST | `/signup` | Create a new user | ❌ Not Required |
| POST | `/login` | Authenticate user & return JWT token | ❌ Not Required |
| POST | `/events/{id}/register` | Register user for an event | ✅ Required |
| DELETE | `/events/{id}/register` | Cancel event registration | ✅ Required |

---

## 📚 Reference
Course: **"Go - The Complete Guide"**  
Instructor: **Maximilian Schwarzmüller**

---
