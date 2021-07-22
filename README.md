# Vegan Recipe Search Engine

This app has scraped data form the web with Go (Golang) to provide a data set of Vegan Recipes. The Website of this app is made with React.js & provides a vegan recipe search engin. I hope you link it!

[Visit Site](https://lwrgithub.github.io/vegan-recipes-search-engine/)


## Tech Used

- JavaScript
- Node.js
- React.js
- Go (Golang)
- HTML/CSS
- Boostrap
- [GitHub Pages](https://lwrgithub.github.io/vegan-recipes-search-engine/)


## Setup, Install & Run

### Install
***Note: make sure to change the files to reflect your GitHub info:*** [Learn How](https://blog.usejournal.com/how-to-deploy-your-react-app-into-github-pages-b2c96292b18e)

```
yarn install
```

**(Not Necessary) Run the Go Web Scraper:**

```
cd go
go run scrape.go
```

### Run
***@ localhost:3000***

```
yarn start
```

### Build/Push
***When using GitHub Pages do a build before pushing***

```
npm run deploy
git add -A
git commit -m "Some New Code"
git push
```

**Troubleshooting Bootstrap Cards**

The curent vertion of Bootstrap as of ***Saterday July 10, 2021*** has some kind of bug in it when it comes to Cards so I have provided an old vertion of `React-Bootstrap` in the file named "bootstrap" 
- Simply replace the bootstrap file in node_modules with the one provided.


## IMG

<img alt="Screenshot of the home page. The image has photos of the Different vegan dishes with their names such as tomato soup & slaw with coriander." src="https://github.com/LWRGitHub/vegan-recipes-search-engine/tree/main/public/img/vegan-search-engine-home.png" width="100%" />
