import { captureRejectionSymbol } from 'events';
import express from 'express';

const app = express();

const fetch = (...args) => 
    import ('node-fetch').then(({default: fetch}) => fetch(...args));
const baseURL = "http://localhost:8080/items";

app.get("/api/items/:id", async function(req, res){
    res.set("Acces-Control-Allow-Origin", "*");
    const url = `${baseURL}/${req.params.id}`;
    const options = {method: 'GET'};
    try{
        let response = await fetch(url, options);
        let status = response.status;
        response = await response.json();
        res.status(status).json(response);
    }catch (err){
        console.log(err);
        res.status(500).json({msg: 'Internal Server Error'});
    }
});
app.listen(5001, () => {
    console.log("Server started on port 5001");
});