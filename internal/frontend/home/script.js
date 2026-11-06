function main() {
  document.getElementById("repo-link-schema").innerText = location.origin + "/r/{user}/{repo_name}";
}

window.addEventListener(
  "DOMContentLoaded",
  main
);
