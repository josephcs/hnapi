### HackerNews API in Go

Had 10 minutes to kill and a simple RESTful API to fetch posts from HackerNews was born. _First page hot posts for now_.

#### Usage

__GET__ `/`

```json
{
    "posts": [
        {
            "title": "Google is acquiring Kaggle",
            "url": "https://techcrunch.com/2017/03/07/google-is-acquiring-data-science-community-kaggle/",
            "story_url": "https://news.ycombinator.com/item?id=13817883",
            "score": 480
        },
        {
            "title": "Attitudes That Lead to Maintainable Code",
            "url": "https://spin.atomicobject.com/2017/03/07/attitudes-maintainable-code/",
            "story_url": "https://news.ycombinator.com/item?id=13819604",
            "score": 24
        }
    ]
}
```
