<p align="center" style="font-size:50px;"> <b> IN PROGRESS </b> </p>
<p align="center">
    <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" align="center" width="30%">
</p>
<p align="center"><h1 align="center">THEWIIDOWNLOADER</h1></p>
<p align="center">
	<em><code>❯ REPLACE-ME</code></em>
</p>
<p align="center">
	<!-- local repository, no metadata badges. --></p>
<p align="center">Built with the tools and technologies:</p>
<p align="center">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=default&logo=Go&logoColor=white" alt="Go">
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

    Downloads Wii Games and format them for usb loader gx

---

##  Features

    1. Grab Links
    2. Download from any link
    3. Unzip
    4. Covert RVZ To ISO
    5. Convert ISO To WBFS format

---

##  Project Structure

```sh
└── thewiidownloader/
    ├── go.mod
    ├── go.sum
    ├── links.txt
    ├── main.go
    ├── modules
    │   ├── config.go
    │   ├── downloader.go
    │   ├── getLinks.go
    │   ├── log.go
    │   ├── remove.go
    │   ├── staus.go
    │   ├── toiso.go
    │   ├── towbfs.go
    │   ├── types.go
    │   └── unzip.go
    ├── README.md
    └── tls
        └── request.go
```


###  Project Index
<details open>
	<summary><b><code>THEWIIDOWNLOADER/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/go.mod'>go.mod</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/go.sum'>go.sum</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/links.txt'>links.txt</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/main.go'>main.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- modules Submodule -->
		<summary><b>modules</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/config.go'>config.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/downloader.go'>downloader.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/getLinks.go'>getLinks.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/log.go'>log.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/remove.go'>remove.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/staus.go'>staus.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/toiso.go'>toiso.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/towbfs.go'>towbfs.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/types.go'>types.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/modules/unzip.go'>unzip.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- tls Submodule -->
		<summary><b>tls</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/P-0001/thewiidownloader/tls/request.go'>request.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with thewiidownloader, ensure your runtime environment meets the following requirements:

- **Programming Language:** GoLang 1.23
- **Dolphin Emulator:** [Dolphin Emulator](https://dolphin-emu.org/)
- **ISO to WBFS:** [ISO to WBFS](https://www.isotowbfs.com/)


###  Installation

Install thewiidownloader using one of the following methods:

**Build from source:**

1. Clone the thewiidownloader repository:
```sh
❯ git clone https://github.com/P-0001/thewiidownloader
```

2. Navigate to the project directory:
```sh
❯ cd thewiidownloader
```

3. Install the project dependencies:

```sh
❯ go mod tidy
```

**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```

<h1> @todo add .env * </h1>


###  Usage
Run thewiidownloader using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run main.go
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```


---
##  Project Roadmap

- [X] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

##  Contributing

- **💬 [Join the Discussions](https://github.com/P-0001/thewiidownloader/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/P-0001/thewiidownloader/issues)**: Submit bugs found or log feature requests for the `thewiidownloader` project.
- **💡 [Submit Pull Requests](https://github.com/P-0001/thewiidownloader/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your LOCAL account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/P-0001/thewiidownloader
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to LOCAL**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com/P-0001/thewiidownloader/graphs/contributors">
      <img src="https://contrib.rocks/image?repo=P-0001/thewiidownloader">
   </a>
</p>
</details>

---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[Dolphin Emulator](https://dolphin-emu.org/)
[ISO to WBFS](https://www.isotowbfs.com/)


---
