let currentPage = 1;
let totalPages = 1;
let word = "";

async function getData() {
    event.preventDefault(); // Prevent form from reloading the page
    
    document.getElementById("content").innerHTML = "";

    let query = document.getElementById("searchInput").value.trim();
    if (query.length === 0) return; // Do nothing if input is empty

    if (word != word) {
        currentPage = 1;
    }

    try {
        let response = await fetch(`/${query}/${currentPage}`);
        if (!response.ok) throw new Error("Failed to fetch data");

        let obj = await response.json();
        data = obj.data;
        console.log(data);
        
        currentPage = data.page;
        totalPages = data.total_pages;
        word = data.word;

        renderData(data, currentPage);
        
    } catch (error) {
        console.error("Error fetching data:", error);
        document.getElementById("content").innerHTML = `<p style="color: red;">Error fetching data.</p>`;
    }
}

function renderData(data, currentPage) {
    const content = document.getElementById("content");
    const ul = document.createElement("ul");

    for (verse of data.results) {
        let li = document.createElement("li");
        li.innerHTML = `${verse.reference}: ${verse.content}`;

        ul.appendChild(li);
    }
    content.appendChild(ul);

    if (currentPage != 1) {
        document.getElementById("prev").disabled = false;
    } else {
        document.getElementById("prev").disabled = true;
    }

    if (currentPage < totalPages) {
        document.getElementById("next").disabled = false;
    } else {
        document.getElementById("next").disabled = true;
    }

    document.getElementById("page-info").innerHTML = `<div>Page ${currentPage}/${totalPages}</div>`;
}

function prevPage() {
    currentPage -= 1;

    getData();
}

function nextPage() {
    currentPage += 1;

    getData();
}
