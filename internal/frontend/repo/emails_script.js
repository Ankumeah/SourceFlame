async function main() {
  const repo = window.location.pathname.slice(3, -6);
  document.title = repo + " emails"

  const url = new URL(window.location.href);
  if (url.searchParams.get("page") === null) {
    url.searchParams.set("page", "1");
    window.location.href = url;
  }

  const page = url.searchParams.get("page");

  loadEmails(repo, page);
  setNextAndLast(page);
}

async function loadEmails(repo, page) {
  const res = await fetch(
    `${location.origin}/r/` +
    `${repo}/emails` +
    `?page=${page}`
  );
  if (!res.ok) {
    throw new Error(`Error loading emails: ${res.status}`);
  }

  let emails = {};
  try {
    emails = await res.json();
  } catch (err) {
    console.log(`Invalid json from emails api: ${err}`);
  }

  const emailsDiv = document.getElementById("emails")

  emails.forEach(
    (email) => {
      let data = atob(email.data || "RW1wdHkK");

      let details = document.createElement("details");
      emailsDiv.appendChild(details);

      let summery = document.createElement("summary");
      details.appendChild(summery);

      let h3 = document.createElement("h3");
      h3.innerText = getSubject(data);
      summery.appendChild(h3);

      let pre = document.createElement("pre");
      pre.className = "code-block";
      pre.innerText = data;
      details.appendChild(pre);
    }
  )
}

function setNextAndLast(page) {
  document.getElementById("next").href = "?page=" + (+page + 1);
  document.getElementById("last").href = "?page=" + Math.max(page - 1, 1);
}

function getSubject(rawEmail) {
  const match = rawEmail.match(/^Subject:(.*)$/m);

  return match ? match[1].trim() : "No subject";
}

window.addEventListener(
  "DOMContentLoaded",
  main
);
