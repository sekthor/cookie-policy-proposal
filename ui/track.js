function getLocation() {
  return new Promise((resolve, reject) => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(resolve, reject);
    } else {
      reject(new Error("Your browser does not have geolocation. We cannot locate you :("));
    }
  });
}

async function success(position) {
  console.log("starting success")
  try {
    const headers = new Headers({ "Content-Type": "application/json" });

    const response = await fetch("/api/location", {
      method: "POST",
      body: JSON.stringify(position),
      headers,
    });
     
    console.log("sending location")

    if (!response.ok) {
      alert("Could not send location");
    }
  } catch (error) {
    console.error("Error sending location:", error);
  }
}

function hasAcceptedCookies() {
  return localStorage.getItem("cookies-accepted") === "true";
}

async function track() {
  if (hasAcceptedCookies()) {
    try {
      const position = await getLocation();
      await success(position);
      window.location = "https://www.google.com/maps/search/bakery";
    } catch (error) {
      alert("Sorry, no position available.");
    }
  }
}