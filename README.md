# Spike

This is intended to answer the [questions outlined here](https://github.com/creativesoftwarefdn/weaviate/issues/794).

The following is copied from the mentioned issue:

---

## Background

Our original plan was to use OIDC as an auth module as outlined in #creativesoftwarefdn/weaviate#628.
In the meantime we have learned that [there is also a way to integrate
oauth2](https://github.com/go-swagger/go-swagger/tree/master/examples/composed-auth)
directly into `go-swagger`.

## Goal

**Question: Would we benefit from using this feature over building the required
things ourselves?**

In more detail:
* Do we lose pluggability because of this?
* How would we deal with the fact that many of our Auth decisions are happening
  in GraphQL which from a REST perspective is always the same endpoint?
* Will we stay flexible enough to support custom enterprise Identity Providers
  and Token Issuers?
* How much effort would we save? Would any additional work be required that
  wouldn't be necessary otherwise?

