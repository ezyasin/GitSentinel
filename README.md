

# GitSentinel

**GitSentinel** is an open-source command-line tool that analyzes GitHub repositories to generate **risk** and **reputation scores**, helping developers and organizations make informed decisions about the code they use or evaluate.

---

## 🚀 Features

* Objective repository analysis
* Risk and reputation scoring
* Insightful metrics on code quality, activity, and security
* Simple and efficient CLI interface

---

## 📦 Installation

GitSentinel requires **Go 1.21+** to build or install.

### Option 1: Install via `go install`

```bash
go install github.com/ezyasin/git-sentinel/cmd/git-sentinel@latest
```

---

## 🧪 Usage

Analyze a GitHub repository:

```bash
git-sentinel analyze https://github.com/owner/repo
```

---

## 📊 Output

GitSentinel provides:

* **Total Score** (0.00 - 1.00)
* **Risk Score** (0.00 - 1.00)
* **Reputation Score** (0.00 - 1.00)
* Repository Metrics:

  * Stars & Forks
  * Open Issues & Pull Requests
  * Last Commit Date
  * README Presence
  * Account Age
  * Follower Count
  * Repository Count

---

## 📈 Scoring Criteria

Scores are based on:

* **Activity**: Commit frequency, issue responsiveness
* **Community Engagement**: Stars, forks, followers
* **Maintenance**: README availability, open/closed issues
* **Account History**: Age of the account, total repositories

---

## 🛠 Development

### Prerequisites

* Go 1.21 or higher
* GitHub API token (optional but recommended for increased rate limits)

### Build Instructions

```bash
git clone https://github.com/ezyasin/git-sentinel.git
cd git-sentinel
go build -o git-sentinel ./cmd/git-sentinel
```

---

## 🤝 Contributing

Contributions are welcome!
Feel free to open issues, suggest features, or submit pull requests.