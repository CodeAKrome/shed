# shed
Toolshed. This repo contains my custom tools and a few others I've copied.
## [media origin](https://github.com/CodeAKrome/bootcupboard/tree/main/llm-test/go-media/media)
List media files in a directory tree. Outputs 2 column list with file type [image,book,text,comic,video,audio] and file path.
### Sample
```sh
media ~/Pictures
```
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

---

## flair-news
Uses [flair](https://github.com/flairNLP/flair) to identify entities (NER) which [NewsMTSC](https://github.com/fhamborg/NewsMTSC?tab=readme-ov-file) uses to perform targetted sentiment analysis.

---

# Citations
## [flair](https://github.com/flairNLP/flair) framework:

```
@inproceedings{akbik2019flair,
  title={{FLAIR}: An easy-to-use framework for state-of-the-art {NLP}},
  author={Akbik, Alan and Bergmann, Tanja and Blythe, Duncan and Rasul, Kashif and Schweter, Stefan and Vollgraf, Roland},
  booktitle={{NAACL} 2019, 2019 Annual Conference of the North American Chapter of the Association for Computational Linguistics (Demonstrations)},
  pages={54--59},
  year={2019}
}

```
## [NewsMTSC](https://github.com/fhamborg/NewsMTSC?tab=readme-ov-file)
### [paper](https://aclanthology.org/2021.eacl-main.142/) ([PDF](https://aclanthology.org/2021.eacl-main.142.pdf)):

```
@InProceedings{Hamborg2021b,
  author    = {Hamborg, Felix and Donnay, Karsten},
  title     = {NewsMTSC: (Multi-)Target-dependent Sentiment Classification in News Articles},
  booktitle = {Proceedings of the 16th Conference of the European Chapter of the Association for Computational Linguistics (EACL 2021)},
  year      = {2021},
  month     = {Apr.},
  location  = {Virtual Event},
}
```