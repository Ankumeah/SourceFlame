# Routes
This file contains all routes in this project
For a more formal spec check `openapi.yaml`

> [!NOTE]
> While this file was hand written by me, the `openapi.yaml`
spec is AI generated with this file as the example


| Method | Route                                               | Queries                                               | Headers                                          | Body              |
|--------|-----------------------------------------------------|-------------------------------------------------------|--------------------------------------------------|-------------------|
| GET    | `/`                                                 |                                                       |                                                  |                   |
| GET    | `/r/{user}/{repo}`                                  |                                                       |                                                  |                   |
| GET    | `/r/{user}/{repo}/info/refs`                        | `service` (`git-receive-pack` \| `git-upload-pack`)   |                                                  |                   |
| POST   | `/r/{user}/{repo}/git-upload-pack`                  |                                                       |                                                  | Git stuff idk     |
| POST   | `/r/{user}/{repo}/git-receive-pack`                 |                                                       | `Authorization (Basic base64(user:pass))`        | Git stuff idk     |
| GET    | `/r/{user}/{repo}/info/refs`       | `service` (`git-receive-pack` \| `git-upload-pack`)   |                                                  |                   |
| POST   | `/r/{user}/{repo}/git-upload-pack` |                                                       |                                                  | Git stuff idk     |
| POST   | `/r/{user}/{repo}/git-receive-pack`|                                                       | `Authorization (Basic base64(user:pass))`        | Git stuff idk     |
| GET    | `/r/{user}/{repo}/refs`            |                                                       |                                                  |                   |
| GET    | `/r/{user}/{repo}/blob/{file...}`  | `ref`                                                 |                                                  |                   |
| GET    | `/r/{user}/{repo}/git-log`         | `ref`                                                 |                                                  |                   |
| GET    | `/r/{user}/{repo}/tree/{dir...}`   | `ref`                                                 |                                                  |                   |
| GET    | `/api/ping`                            |                                                       |                                                  |                   |
| POST   | `/api/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| PATCH  | `/api/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        | `{ "email": "" }` |
| DELETE | `/api/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| POST   | `/api/repos/{repo}`                    |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| DELETE | `/api/repos/{repo}`                    |                                                       | `Authorization (Basic base64(user:pass))`        |                   |

