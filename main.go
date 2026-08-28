package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// Chapter 是章节列表的一项
type Chapter struct {
	File  string `json:"file"`  // 文件名，如 01-安装与环境搭建.md
	Title string `json:"title"` // 显示标题（去掉序号前缀和 .md 后缀）
}

// listChapters 扫描 chapters/ 目录，返回按文件名排序的章节列表
func listChapters() ([]Chapter, error) {
	entries, err := os.ReadDir("chapters")
	if err != nil {
		return nil, err
	}

	chapters := make([]Chapter, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		// "01-安装与环境搭建" → 标题去掉序号前缀，侧边栏显示更干净
		title := strings.TrimSuffix(e.Name(), ".md")
		if idx := strings.Index(title, "-"); idx != -1 {
			title = title[idx+1:]
		}
		chapters = append(chapters, Chapter{
			File:  e.Name(),
			Title: title,
		})
	}

	sort.Slice(chapters, func(i, j int) bool {
		return chapters[i].File < chapters[j].File
	})
	return chapters, nil
}

func main() {
	// 静态资源（style.css / script.js）
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 章节 Markdown 原文（前端用 marked.js 渲染成 HTML）
	http.Handle("/chapters/", http.StripPrefix("/chapters/", http.FileServer(http.Dir("chapters"))))

	// 章节列表 API：返回 chapters/ 下的所有 .md 文件名
	http.HandleFunc("/api/chapters", func(w http.ResponseWriter, r *http.Request) {
		chapters, err := listChapters()
		if err != nil {
			http.Error(w, "list chapters failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chapters)
	})

	// 首页：返回页面骨架模板
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "courses/layout.html")
	})

	addr := ":8080"
	log.Printf("Go 学习站点启动: http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
