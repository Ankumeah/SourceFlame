async function main() {
  const repo = window.location.pathname.slice(3, -6);
  document.title = repo + " emails"

  loadEmails(repo);
}

async function loadEmails(repo) {
  const res = await fetch(
    `${location.origin}/r/` +
    `${repo}/emails`
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

function getSubject(rawEmail) {
  const match = rawEmail.match(/^Subject:(.*)$/m);

  return match ? match[1].trim() : "No subject";
}

window.addEventListener(
  "DOMContentLoaded",
  main
);
