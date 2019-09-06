---
id: "spec_changelog"
title: "twirk Wire Protocol Changelog"
sidebar_label: "Changelog"
---

This document lists changes to the twirk specification.

## Changed in v6

### URL scheme

In [v5](./PROTOCOL.md), URLs followed this format:

```bnf
**URL ::= Base-URL "/twirk/" [ Package "." ] Service "/" Method**
```

Version 6 changes this format to remove the mandatory `"/twirk"` prefix:

```bnf
**URL ::= Base-URL "/" [ Package "." ] Service "/" Method**
```

Also, `Base-URL` can now contain a path component - in other words, it's legal
to set any prefix you like.

If you loved the old `/twirk` prefix, you can still use it by using a base URL
that ends with `/twirk`. You're no longer forced into use it, however.

The "/twirk/" prefix is no longer required for three reasons:

 - Trademark concerns: some very large organizations don't want to
   take any legal risks and are concerned that "twirk" could become
   trademarked.
 - Feels like advertising: To some users, putting "twirk" in all your
   routes feels like it's just supposed to pump twirk's brand, and
   provides no value back to users.
 - Homophonous with "twerp": In some Very Serious settings (like
   government websites), it's not okay that "twirk" sounds like
   "twerp", which means something like "insignificant pest."
