# SourceFlame

SourceFlame is a minimal git server written in [Golong](https://go.dev)

> [!NOTE]
> This is a rewrite of [SourceFlameBackend](https://github.com/Ankumeah/SourceFlameBackend).
After looking back at it, it swayed too much from its goal of minimalism and I
decided to do a from scratch rewrite with all my new knowledge of softwear in hopes
of better design decisions

## How to run
> [!NOTE]
> Make sure to check and edit env.example to your use

```sh
go build -o build/sf -ldflags='-s -w' .

export $(cat env.example)
mkdir -p build/repos

./build/sf
```

or

```sh
docker build -t sf .

docker run \
  --name sf \
  --env-file env.example \
  --volume ./data:/opt/sf/build/repos \
  sf
```

## How to use
Take a look at `test.d/` and you should get the general idea (actual docs soon)
