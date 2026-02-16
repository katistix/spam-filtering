# spam-filtering - implementing Naive Bayes Theorem in GO

Heavily inspired by [Tsoding's "Email Spam Filter in Go" video](https://www.youtube.com/watch?v=JsfOXk7qmSM).


Formula:

$$
P(C|D) = \frac{P(D|C) \cdot P(C)}{P(D)}
$$

where: 
- $P(D)$ - probability of a given document to be exist based on the current Bag-of-Words (order of words is not relevant)
- $P(C)$ - probability of a `class` of document to be `True` (in this case the `SPAM class`)


## # get the dataset

Go to https://www2.aueb.gr/users/ion/data/enron-spam/ and download and extract `enron1`, `enron2` etc. and put them in `data/`



## Progress:
- [x] generate Bag-of-Words for a directory
- [x] compute $P(D)$ for a given document (at a `filePath`)
- [ ] compute $P(C)$