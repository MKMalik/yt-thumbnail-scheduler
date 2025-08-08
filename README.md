# 📸 YouTube Thumbnail Scheduler — `ThumbnailIntel.ai`

An intelligent, automated thumbnail optimization SaaS tool built for YouTube creators.  
It dynamically swaps thumbnails based on live performance data like views, likes, top comments, any events/festivals and more — boosting CTR and streamlining creative workflows.

> **Set it. Forget it. Grow your channel.**

> **Clickbait? Nah. Smart Thumbnail.**

---

## 🚀 Features

- ✅ **Rule-based Thumbnail Switching**
  - Example: *"If views > 10k and likes < 500 → Switch to AltImage2.png"*

- 🔁 **A/B/C Thumbnail Rotation**
  - Rotate thumbnails across defined intervals
  - Track engagement for each

- 💬 **Comment-Aware Overlay Engine**
  - Pull top-liked comments
  - Auto-overlay onto thumbnails

- 🕘 **Scheduled Campaigns**
  - Day-wise rotation (e.g., Day 1 → AltImage1, Day 3 → AltImage2, revert on Day 5)

- 📊 **Analytics Dashboard**
  - Visual CTR comparison after each switch
  - Engagement heatmaps

- ⚠️ **Auto-Revert System**
  - Automatically revert to best-performing thumbnail if metrics drop

---

## 🧱 System Architecture

![System Architecture](./idea/ThumbnailIntelli.png)  


**Core Components:**
- Thumbnail Rules Engine
- Metric Poller (YouTube Data API)
- Thumbnail Action Scheduler
- Queue + Worker (Redis/BullMQ)
- Overlay Renderer (Optional Puppeteer or Canvas)
- Admin UI (HTML/CSS/JS or Next.js)

---

## 🧰 Tech Stack

| Layer        | Tech                        |
|--------------|-----------------------------|
| Backend      | **Go** (or Node.js for MVP) |
| Queue        | **BullMQ** (Redis)          |
| Database     | **MySQL**                   |
| Frontend     | HTML/CSS/JS *(MVP)*         |
| Infra        | Docker, PM2, GitHub CI/CD   |
| APIs         | YouTube Data API v3         |
| Extras       | Puppeteer (for overlays)    |

---

## 🧠 How It Works

1. Creator connects YouTube channel
2. Defines rules and uploads thumbnails
3. Worker polls metrics (views, likes, comments)
4. Rules engine triggers thumbnail switch (if matched)
5. System tracks post-switch performance
6. Auto revert if performance drops
7. Optional: Best comment is overlaid on thumbnail

---
