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
| GET    | `/api/$API_VERSION/ping`                            |                                                       |                                                  |                   |
| POST   | `/api/$API_VERSION/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| PATCH  | `/api/$API_VERSION/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        | `{ "email": "" }` |
| DELETE | `/api/$API_VERSION/acc`                             |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| GET    | `/api/$API_VERSION/r/{user}/{repo}/info/refs`       | `service` (`git-receive-pack` \| `git-upload-pack`)   |                                                  |                   |
| POST   | `/api/$API_VERSION/r/{user}/{repo}/git-upload-pack` |                                                       |                                                  | Git stuff idk     |
| POST   | `/api/$API_VERSION/r/{user}/{repo}/git-receive-pack`|                                                       | `Authorization (Basic base64(user:pass))`        | Git stuff idk     |
| GET    | `/api/$API_VERSION/r/{user}/{repo}/refs`            |                                                       |                                                  |                   |
| GET    | `/api/$API_VERSION/r/{user}/{repo}/blob/{file...}`  | `ref`                                                 |                                                  |                   |
| GET    | `/api/$API_VERSION/r/{user}/{repo}/git-log`         | `ref`                                                 |                                                  |                   |
| GET    | `/api/$API_VERSION/r/{user}/{repo}/tree/{dir...}`   | `ref`                                                 |                                                  |                   |
| POST   | `/api/$API_VERSION/repos/{repo}`                    |                                                       | `Authorization (Basic base64(user:pass))`        |                   |
| DELETE | `/api/$API_VERSION/repos/{repo}`                    |                                                       | `Authorization (Basic base64(user:pass))`        |                   |

