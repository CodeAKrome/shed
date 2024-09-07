# shed
Toolshed. This repo contains my custom tools and a few others I've copied.
## [media origin](https://github.com/CodeAKrome/bootcupboard/tree/main/llm-test/go-media/media)
List media files in a directory tree. Outputs 2 column list with file type [image,book,text,comic,video,audio] and file path.
### Sample
```
image	/Users/kyle/Pictures/Pile/10meter.png
book	/Users/kyle/Pictures/Pile/Band Chart - 11X17 Color.pdf
image	/Users/kyle/Pictures/Pile/Husky.jpeg
video	/Users/kyle/Pictures/Pile/IMG_0548.MOV
-- snip --

Summary:
book: 10
image: 549
video: 12
Total files found: 571
Errors encountered: 0
Depth: Unlimited
Error limit: Unlimited
Media types searched: [image book text comic video audio]
```

---

## zsh Z-shell functions and things to go into .zshrc type places
### fzf_history.zsh
ctrl-r allows you to fuzzy search your command history using fzf