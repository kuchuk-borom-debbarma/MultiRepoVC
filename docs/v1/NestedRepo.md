Below is a **complete, clean, production-ready Markdown design document** for your **Nested Repository Implementation** in MRVC.
You can paste this directly into your project docs.

---

# 📘 MRVC Nested Repository Design Document

## Overview

Nested repositories are a core feature of MRVC that allow independent repositories to exist inside other repositories.
A nested repo operates **fully independently**, while the parent repo maintains **snapshot-based references** to it.

This document defines:

* How nested repos are detected
* How they are represented as objects
* How they interact with commit snapshots
* How renames/moves are handled
* Rules for repository uniqueness
* How nested repos integrate into MRVC’s content-addressable storage model

---

# 🚀 Goals of Nested Repo Support

1. **Independence**
   Each repo has its own `.mrvc` folder, metadata, objects, HEAD, commits.

2. **Deterministic Snapshotting**
   Parent commits store *what nested repos existed at that moment*.

3. **Robustness Against Moves/Renames**
   Path changes between commits should be detected naturally.

4. **Consistency With MRVC’s Design**
   Everything must be an **object** and **content-addressable**.

5. **Zero Manual Maintenance**
   Snapshot metadata should always auto-adjust.

---

# 🧩 Core Concepts

## 1. What is a Nested Repo?

A nested repo is any directory that contains a **`.mrvc` folder** inside a parent MRVC repo.

Example:

```
parent-repo/
  .mrvc/
  services/
    auth-service/
      .mrvc/
```

`auth-service` is a nested MRVC repository.

---

## 2. Unique Repository Names (Required)

To avoid ambiguity:

* Every repository must have a **unique name** within its entire ancestor chain.
* No child repo may share the same name as any parent or sibling repo.

**Validation during `mrvc init`:**

1. Check parent directories for `.mrvc`.
2. Load metadata from each parent.
3. Ensure the new repo name does NOT match any ancestor name.

This ensures stable identification and avoids collisions.

---

## 3. Repository Identity (`repo_id`)

Each repo has a **stable, immutable identity** generated at initialization:

* Preferably a **UUID**
* Stored in `.mrvc/metadata.json`

Example:

```json
{
  "name": "auth-service",
  "author": "kuchuk",
  "created_at": "123456789",
  "repo_id": "8cf2d94a-8d3f-45fb-a694-44dcd50917da"
}
```

`repo_id` never changes unless this is a different repo.

---

# 🗂️ Nested Repo Object

Nested repos are stored as **first-class objects** inside `.mrvc/objects`.

This is MRVC’s 4th object type:

* **blob** — file content
* **tree** — directory structure
* **commit** — snapshot
* **nested_repo** — pointer to another repo

---

## NestedRepoObject Structure

```json
{
  "repo_id": "uuid-123",
  "name": "auth-service",
  "path": "services/auth"
}
```

This JSON describes:

* **repo_id** → permanent identity
* **name** → unique name within hierarchy
* **path** → location at the time of the commit

---

## NestedRepoObject Hashing

Hash = SHA256(JSON)

Hash stability rules:

* **Stays the same** when nested repo files/commits change
* **Changes** only if:

    * the repo **moves**
    * the repo **renames**
    * or the repo_id changes (new repo initialized)

This perfectly matches snapshot behavior.

---

## Storage

Stored in MRVC object database:

```
.mrvc/objects/<first2>/<rest>
```

Example:

```
.mrvc/objects/a1/b2c3d4e5f6....
```

---

# 🧱 Commit Object: Nested Repos

Commit object is extended with a new field:

```json
{
  "tree": "...",
  "parent": "...",
  "message": "...",
  "author": "...",
  "timestamp": "...",
  "nested_repos": [
    "a1b2c3d4e5...",
    "ffeedd9933..."
  ]
}
```

Only the **hashes** of NestedRepoObjects are stored.

Not the objects themselves.

This follows MRVC’s content-addressable design:

* Trees reference blob/tree hashes
* Commits reference tree hashes
* Commits reference nested repo hashes

Everything is a pointer.

---

# 🔍 Directory Scan Rules

During commit:

1. Use filesystem walking with configurable options:

    * Skip `.mrvc` folders
    * Skip nested repos (folders with their own `.mrvc`)

2. When nested repo detected:

    * Read `metadata.json`
    * Construct `NestedRepoObject`
    * Hash it
    * Save the object
    * Add hash to commit’s `nested_repos` array

3. Continue scanning normally for other files.

---

# 🔄 Snapshot Semantics

### Example timeline:

Commit #1:

```
services/auth
```

Commit #2 (repo moved):

```
modules/authentication
```

Commit #1 nested object:

```json
{
  "repo_id": "uuid-123",
  "name": "auth-service",
  "path": "services/auth"
}
```

Commit #2 nested object:

```json
{
  "repo_id": "uuid-123",
  "name": "auth-service",
  "path": "modules/authentication"
}
```

Hash changes.
Snapshot reflects rename/move.

### Child repo commits do NOT require parent commit updates.

Their content doesn't affect the parent.

---

# 🔐 Reconstruction Use Case

When checking out a commit:

1. Load commit object.
2. Restore tree structure.
3. Load nested repo hashes:

    * read NestedRepoObject
    * ensure path exists
    * (optional future feature) reattach nested repo clone

This enables a future “super checkout” behavior.

---

# 🛡️ Integrity Guarantees

* Parent cannot accidentally include child repo’s files.

* Nested repos remain independent.

* Snapshot captures only:

    * identity
    * name
    * path

* Path changes produce new nested objects.

* Repo moves/renames are automatically reflected.

* Repo contents never affect nested repo object hashes.

---

# 🧠 Benefits of This Design

### ✔ Clean and simple

No `.mrvcmodules` like Git.
No complex coupling.

### ✔ Fully snapshot-consistent

Everything is immutable and content-addressable.

### ✔ Extremely robust

Handles moves, renames, reorganizing directories.

### ✔ Zero user maintenance

Snapshots self-update naturally.

### ✔ Perfect alignment with MRVC philosophy

**“Everything is an object.”**

---

# 📦 Summary

| Concept                 | Behavior                                    |
| ----------------------- | ------------------------------------------- |
| Nested repo detection   | Directory containing `.mrvc`                |
| Nested repo identity    | `repo_id` (UUID)                            |
| Nested repo object      | JSON file hashed and stored in object store |
| Commit representation   | `nested_repos: [hash, hash, ...]`           |
| Move/rename behavior    | New object created, new hash                |
| Child repo independence | Always preserved                            |
| Name uniqueness         | Enforced across ancestor hierarchy          |

TODO
- start commit command :- This will allow step by step commit per nested repo.
- refactor commit command to support nested repo commits :- This will allow specifying repo with ids and messages for commit.