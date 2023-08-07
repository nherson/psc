import axios from "axios";
import { baseUrl } from "../api/client";

export const downloadStats = async () => {
  axios({
    method: "GET",
    url: `https://poopswagchampionship.com/data/upcoming`,
  }).then((resp) => {
    const url = window.URL.createObjectURL(new Blob([resp.data]));
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `psc-stats.csv`);
    document.body.appendChild(link);
    link.click();
    link.parentNode?.removeChild(link);
  });
};
