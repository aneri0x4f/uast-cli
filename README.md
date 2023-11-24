# UAST: Unicode Aware Sanskrit Transliteration

A tool for transliterating and typing Sanskrit in the easiest
computer, language, and human-friendly way

To install this program, you will need to install [Go](https://go.dev/)

After that, simply run:

```bash
go install https://github.com/aneri0x4f/uast-cli/cmd/uast@latest
```

To remove,

```bash
make clean
```

To build in current directory,

```bash
make
```

To run in [Docker](https://www.docker.com/) container,

```bash
make docker
docker run -it --rm uast
```

```
Usage of uast:
  -from string
    	from schema ([uast raw devanagari iast]) (default "uast")
  -i string
    	Input file
  -o string
    	Output file
  -to string
    	to schema ([uast raw devanagari iast]) (default "devanagari")
```

If you use this repository, please cite the following paper:

```bibtex
@misc{uast_2022,
  doi = {10.48550/ARXIV.2203.14277},
  url = {https://arxiv.org/abs/2203.14277},
  author = {Dalwadi, Aneri and Dave, Dhruvil},
  keywords = {Human-Computer Interaction (cs.HC), FOS: Computer and information sciences, H.5.2},
  title = {UAST: Unicode Aware Sanskrit Transliteration},
  publisher = {arXiv},
  year = {2022},
  copyright = {Creative Commons Attribution 4.0 International}
}
```

![poster](/uast_poster.png)

Devanagari (IAST: _devanāgarī_) is the writing system that is adapted by various languages
like Sanskrit (IAST: _saṃskṛta_). International Alphabet of Sanskrit Transliteration (IAST) is
a transliteration scheme for romanisation of Sanskrit language. IAST makes use of diacritics
to represent various characters. On a computer, these are represented using Unicode standard
which differs from how the Sanskrit language behaves at a very fundamental level. This
results in an issue that is encountered while designing typesetting software for _devanāgarī_
and IAST. We hereby discuss the problems and provide a solution that solves the issue of
incompatibilities between various transliteration and encoding schemes.

Web version URL: https://uast.dev

Click on the encoding name for available options.
