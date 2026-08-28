// 加载章节列表到侧边栏
async function loadChapters() {
    const res = await fetch('/api/chapters');
    const chapters = await res.json();

    const nav = document.getElementById('chapter-list');
    nav.innerHTML = '';
    for (const ch of chapters) {
        const a = document.createElement('a');
        a.textContent = ch.title;
        a.href = '#';
        a.onclick = (e) => { e.preventDefault(); loadChapter(ch.file); };
        nav.appendChild(a);
    }
}

// 加载并渲染一章 Markdown
async function loadChapter(file) {
    const res = await fetch('/chapters/' + encodeURIComponent(file));
    if (!res.ok) {
        document.getElementById('content').textContent = '加载失败: ' + res.status;
        return;
    }
    const md = await res.text();
    document.getElementById('content').innerHTML = marked.parse(md);
}

loadChapters();
