# 🐱 Vet Clinic API

Une API REST complète et sécurisée pour la gestion d'une clinique vétérinaire spécialisée dans les chats, développée avec Go, Chi Router et GORM.

## 📋 Table des matières

- [Fonctionnalités](#fonctionnalités)
- [Technologies](#technologies)
- [Installation](#installation)
- [Configuration](#configuration)
- [Utilisation](#utilisation)
- [Documentation API](#documentation-api)
- [Authentification](#authentification)
- [Endpoints](#endpoints)
- [Structure du projet](#structure-du-projet)

## ✨ Fonctionnalités

- **Authentification JWT** : Système de connexion sécurisé avec tokens JWT
- **Gestion des rôles** : Contrôle d'accès basé sur les rôles (admin, user)
- **Gestion des utilisateurs** : CRUD complet pour les comptes utilisateurs
- **Gestion des chats** : CRUD complet pour les profils de chats (nom, âge, race, poids)
- **Gestion des visites** : Suivi des consultations vétérinaires avec date, motif et vétérinaire
- **Gestion des traitements** : Enregistrement et suivi des traitements administrés
- **Historique médical** : Consultation de l'historique complet des visites par chat
- **Filtrage des visites** : Recherche de visites par vétérinaire
- **Documentation Swagger** : Interface interactive pour tester l'API
- **Base de données SQLite** : Stockage persistant avec GORM

## 🛠️ Technologies

- **Go** 1.25.3
- **Chi Router** v5.2.3 - Routeur HTTP léger et performant
- **GORM** v1.31.1 - ORM pour Go
- **SQLite** - Base de données embarquée
- **JWT (golang-jwt)** v4.5.2 - Authentification sécurisée
- **bcrypt** - Hashage des mots de passe
- **Swagger** v1.16.6 - Documentation API automatique
- **http-swagger** v1.3.4 - Interface UI pour Swagger

## 📦 Installation

### Prérequis

- Go 1.25 ou supérieur
- Git

### Étapes d'installation

1. **Cloner le repository**
```bash
git clone https://github.com/emmanuelYohore/vet-clinic-api.git
cd vet-clinic-api
```

2. **Installer les dépendances**
```bash
go mod download
```
swag init
```

4. **Lancer l'application**
```

3. **Lancer l'application**
```bash
go run main.go
```

L'API sera accessible sur `http://localhost:8080`

## ⚙️ Configuration

L'application utilise SQLite comme base de données par défaut. Le fichier `data.db` sera créé automatiquement au premier lancement.

La configuration se trouve dans le package `config` et initialise :
- La connexion à la base de données
- Les repositories pour chaque entité
- Les migrations automatiques des schémas

## 🚀 Utilisation

### Démarrer le serveur

```bash
go run main.go
```

Le serveur démarre sur le port **8080** par défaut.

### Accéder à la documentation Swagger

Une fois le serveur démarré, accédez à l'interface Swagger :

```
http://localhost:8080/swagger/index.html
```

## 📚 Documentation API

L'API utilise Swagger/OpenAPI pour la documentation. Tous les endpoints sont documentés avec :
- Description de la fonctionnalité
- Paramètres requis
- Format des requêtes et réponses
- Codes de statut HTTP

## 🔐 Authentification

L'API utilise JWT (JSON Web Tokens) pour l'authentification. La plupart des endpoints nécessitent un token valide.

### Obtenir un token

**Endpoint** : `POST /login`

**Corps de la requête** :
```json
{
  "username": "admin",
  "password": "password123"
}
```

**Réponse** :
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Utiliser le token

Incluez le token dans l'en-tête de vos requêtes :
```
Authorization: Bearer <votre_token>
```

### Rôles et permissions

- **admin** : Accès complet (création, modification, suppression)
- **user** : Accès lecture seule (consultation des données)

## 🔗 Endpoints

### Authentification (`/login`)

| Méthode | Endpoint | Description | Authentification |
|---------|----------|-------------|------------------|
| `POST` | `/login` | Se connecter et obtenir un token JWT | Non |

### Utilisateurs (`/api/v1/users`)

| Méthode | Endpoint | Description | Rôle requis |
|---------|----------|-------------|-------------|
| `POST` | `/api/v1/users` | Créer un nouvel utilisateur | admin |
| `GET` | `/api/v1/users` | Récupérer tous les utilisateurs | admin, user |
| `GET` | `/api/v1/users/{id}` | Récupérer un utilisateur par ID | admin, user |
| `PUT` | `/api/v1/users/{id}` | Mettre à jour un utilisateur | admin |
| `DELETE` | `/api/v1/users/{id}` | Supprimer un utilisateur | admin |

**Exemple de requête POST** :
```json
{
  "username": "johndoe",
  "password": "securepassword",
  "role": "user"
}
```

### Chats (`/api/v1/cats`)

| Méthode | Endpoint | Description | Rôle requis |
|---------|----------|-------------|-------------|
| `POST` | `/api/v1/cats` | Créer un nouveau chat | admin |
| `GET` | `/api/v1/cats` | Récupérer tous les chats | admin, user |
| `GET` | `/api/v1/cats/{id}` | Récupérer un chat par ID | admin, user |
| `PUT` | `/api/v1/cats/{id}` | Mettre à jour un chat | admin |
| `DELETE` | `/api/v1/cats/{id}` | Supprimer un chat | admin |
| `GET` | `/api/v1/cats/{id}/history` | Récupérer l'historique des visites d'un chat | admin, user |

**Exemple de requête POST** :
```json
{
  "name": "Minou",
  "age": 3,
  "breed": "Persan",
  "weigth": 4500
}
```

### Visites (`/api/v1/visits`)

| Méthode | Endpoint | Description | Rôle requis |
|---------|----------|-------------|-------------|
| `POST` | `/api/v1/visits` | Créer une nouvelle visite | admin |
| `GET` | `/api/v1/visits` | Récupérer toutes les visites | admin, user |
| `GET` | `/api/v1/visits/{id}` | Récupérer une visite par ID | admin, user |
| `PUT` | `/api/v1/visits/{id}` | Mettre à jour une visite | admin |
| `DELETE` | `/api/v1/visits/{id}` | Supprimer une visite | admin |
| `GET` | `/api/v1/cats/{id}/visits` | Récupérer les visites d'un chat | admin, user |
| `GET` | `/api/v1/visits/filter` | Filtrer les visites par vétérinaire | admin, user |

**Exemple de requête POST** :
```json
{
  "date": "2025-12-04T10:30:00Z",
  "motif": "Vaccination annuelle",
  "veterinaire": "Dr. Dupont"
}
```

### Traitements (`/api/v1/treatments`)

| Méthode | Endpoint | Description | Rôle requis |
|---------|----------|-------------|-------------|
| `POST` | `/api/v1/treatments` | Créer un nouveau traitement | admin |
| `GET` | `/api/v1/treatments` | Récupérer tous les traitements | admin, user |
| `GET` | `/api/v1/treatments/{id}` | Récupérer un traitement par ID | admin, user |
| `PUT` | `/api/v1/treatments/{id}` | Mettre à jour un traitement | admin |
| `DELETE` | `/api/v1/treatments/{id}` | Supprimer un traitement | admin |
| `GET` | `/api/v1/visits/{id}/treatments` | Récupérer les traitements d'une visite | admin, user |

**Exemple de requête POST** :
```json
{
  "name": "Antiparasitaire"
}
```

## 📁 Structure du projet

```
vet-clinic-api/
├── main.go                    # Point d'entrée de l'application
├── go.mod                     # Dépendances Go
├── README.md                  # Documentation
├── config/                    # Configuration de l'application
│   └── config.go
├── database/                  # Gestion de la base de données
│   ├── database.go
│   └── dbmodel/              # Modèles de base de données
│       ├── cat.go
│       ├── user.go
│       ├── treatment.go
│       └── visit.go
├── docs/                      # Documentation Swagger générée
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── pkg/                       # Packages applicatifs
    ├── authentification/     # Module d'authentification
    │   ├── controller.go
    │   ├── jwt.go
    │   ├── middleware.go
    │   └── routes.go
    ├── models/               # Modèles de requête/réponse
    │   ├── cat.go
    │   ├── user.go
    │   ├── treatment.go
    │   └── visit.go
    ├── user/                 # Module utilisateurs
    │   ├── controller.go
    │   └── route.go
    ├── cat/                  # Module chats
    │   ├── controller.go
    │   └── routes.go
    ├── visit/                # Module visites
    │   ├── controller.go
    │   └── route.go
    └── treatment/            # Module traitements
        ├── controller.go
        └── route.go
```



## 👤 Auteur

**Emmanuel Yohore**
- GitHub: [@emmanuelYohore](https://github.com/emmanuelYohore)

