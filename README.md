<a name="readme-top"></a>

<div align="center">
  <img src="https://github.com/microverseinc/readme-template/raw/master/murple_logo.png" alt="Microverse Logo" width="140" />
  <br/>
  <h3><b>Quadchecker</b></h3>
</div>

---

# 📗 Table of Contents

- [📖 About the Project](#about-project)
  - [🛠 Built With](#built-with)
    - [Tech Stack](#tech-stack)
    - [Key Features](#key-features)
- [💻 Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Install](#install)
  - [Usage](#usage)
- [👥 Authors](#authors)
- [🔭 Future Features](#future-features)
- [🤝 Contributing](#contributing)
- [⭐️ Show Your Support](#support)
- [🙏 Acknowledgements](#acknowledgements)
- [❓ FAQ](#faq)
- [📝 License](#license)

---

# 📖 Quadchecker <a name="about-project"></a>

> **Quadchecker** is a command-line Go program that reads a shape rendered to stdout and identifies which `quad` function(s) could have produced it — reporting every matching function name and its dimensions.

It reads from **stdin** (fully pipe-friendly), reconstructs the output of each known quad renderer (QuadA–QuadE) for the detected dimensions, and compares. Matches are printed alphabetically and separated by `||`. If nothing matches, it prints `Not a quad function`.

```sh
$ ./quadA 5 3 | ./quadchecker
[quadA] [5] [3]

$ ./quadC 1 1 | ./quadchecker
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

$ echo "hello" | ./quadchecker
Not a quad function
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 🛠 Built With <a name="built-with"></a>

### Tech Stack <a name="tech-stack"></a>

<details>
  <summary>Language</summary>
  <ul>
    <li><a href="https://go.dev/">Go (Golang) 1.21+</a></li>
  </ul>
</details>

<details>
  <summary>Libraries</summary>
  <ul>
    <li>Standard library only — zero external dependencies</li>
    <li><code>bufio</code> — stdin line scanning</li>
    <li><code>sort</code> — alphabetical ordering of matches</li>
    <li><code>strings</code> — pattern generation and comparison</li>
    <li><code>fmt</code> — formatted output</li>
  </ul>
</details>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Key Features <a name="key-features"></a>

- **Identifies all five quad variants** — QuadA through QuadE, each with unique corner characters and border logic
- **Handles multiple matches** — when an input is ambiguous (e.g. a `1×1` `"A"`), every matching quad is printed alphabetically and separated by `||`
- **Pipe-friendly CLI** — reads directly from stdin, composable with any quad binary via the shell pipe operator
- **Strict input validation** — ragged rows (unequal widths) or unrecognised patterns always yield `Not a quad function`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 💻 Getting Started <a name="getting-started"></a>

To get a local copy up and running, follow these steps.

### Prerequisites <a name="prerequisites"></a>

- [Go 1.21+](https://go.dev/dl/) installed on your machine

```sh
go version   # should print go1.21 or higher
```

### Setup <a name="setup"></a>

Clone this repository to your desired folder:

```sh
git clone https://github.com/<your-username>/quadchecker.git
cd quadchecker
```

### Install <a name="install"></a>

Build the binary:

```sh
go build -o quadchecker .
```

### Usage <a name="usage"></a>

Pipe the output of any quad binary into `quadchecker`:

```sh
# Single match
$ ./quadA 5 3 | ./quadchecker
[quadA] [5] [3]

# Multiple matches — ambiguous 1x1 shape
$ ./quadC 1 1 | ./quadchecker
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

# Multiple matches — 1x2 shape
$ ./quadE 1 2 | ./quadchecker
[quadC] [1] [2] || [quadE] [1] [2]

# Not a quad
$ echo "random text" | ./quadchecker
Not a quad function
```

**Quad character reference:**

| Function | Top-Left | Top-Right | Bot-Left | Bot-Right | Edge / Side | Interior |
|----------|:--------:|:---------:|:--------:|:---------:|:-----------:|:--------:|
| QuadA    | `o`      | `o`       | `o`      | `o`       | `-` / `\|`  | ` `      |
| QuadB    | `/`      | `\`       | `\`      | `/`       | `*`         | ` `      |
| QuadC    | `A`      | `A`       | `C`      | `C`       | `B`         | ` `      |
| QuadD    | `A`      | `C`       | `A`      | `C`       | `B`         | ` `      |
| QuadE    | `A`      | `C`       | `C`      | `A`       | `B`         | ` `      |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 👥 Authors <a name="authors"></a>

👤 **Your Name**

- GitHub: [@your-github-handle](https://github.com/your-github-handle)
- Twitter: [@your-twitter-handle](https://twitter.com/your-twitter-handle)
- LinkedIn: [your-linkedin](https://linkedin.com/in/your-linkedin-handle)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 🔭 Future Features <a name="future-features"></a>

- **Support for additional quad variants** — extend the checker as new quad functions (QuadF, QuadG, …) are introduced
- **Reverse mode** — given a function name and dimensions, print what the output would look like
- **`--json` output flag** — machine-readable match results for scripting and tooling

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 🤝 Contributing <a name="contributing"></a>

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](../../issues/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## ⭐️ Show Your Support <a name="support"></a>

If you found this project useful, give it a ⭐️ — it helps others discover it too!

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 🙏 Acknowledgements <a name="acknowledgements"></a>

- [Microverse](https://www.microverse.org/) — for the README template and project structure guidelines
- [Zone01 / 01Founders](https://01founders.co/) — for the original `quad` function exercise that inspired this tool

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## ❓ FAQ <a name="faq"></a>

- **Why does my input match multiple quads?**

  Some quad shapes are geometrically identical at certain dimensions. For example, QuadC, QuadD, and QuadE all render a single `A` character for a `1×1` input, so all three are reported. The checker correctly lists every valid match.

- **What happens if I pass empty input or rows of unequal width?**

  Both cases print `Not a quad function`. The checker requires all rows to share the same width to be considered a valid rectangular shape.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p># quadchecker
