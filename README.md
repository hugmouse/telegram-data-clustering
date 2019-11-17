# Telegram Data Clustering contest 

*This repository is a joke.* 

## Content

-----

- [Telegram Data Clustering contest](#telegram-data-clustering-contest)
  * [Demo](#demo)
  * [Video demo](#video-demo)
  * [Current problem](#current-problem)
  * [Installing and using](#you-are-crazy-and-want-to-try-it-yourself)
  
-----

Here I played with the data that Telegram provided during the clustering contest.

There is too much data to host on a free plan of [zeit](https://zeit.co), so in fact this repository cannot be used for visual representation of all this data. (And zeit.co incorrectly displayed it. Therefore, if you want to test - it is better to go to the installation section)

However, for all who looked, I parsed 166K files `.html` and made in a separate repository - [/tdc-exported/](https://github.com/hugmouse/tdc-exported)

## Demo

![demo](https://i.imgur.com/LtqNv4Y.jpg)

## Video demo

![gifdemo](https://s5.gifyu.com/images/conv.gif)

MP4: https://gfycat.com/immediatepettyappaloosa

## Current problem

`Hugo` cannot process more than 50,000 pages. It just stops and consumes memory and CPU endlessly.

At >20,000 pages, it also stops opening in the web browser.

At this point in time, the maximum number of rendered pages that open in the browser reaches 16,000.

## You are crazy and want to try it yourself

1. First, copy the repository:
```
git clone https://github.com/hugmouse/telegram-data-clustering.git
```

2. Create this path: `./content/posts/`

3. Open the `/golang-related/htmlToMD/recursion.go` and change `consts` for your own path.

```
dirname    = "absolute path to .html data"
dirnameDestination = "absolute path for generated .md data"
```

4. Run `recursion.go` and wait!

*On Linux, there is a memory leak.*
*On Windows, this does not happen.*

5. When markdown is generated you can start the `Hugo server`.
