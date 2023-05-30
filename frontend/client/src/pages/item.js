import React, {useEffect, useState} from "react";

import {useParams} from "react-router-dom";

function Item() {
    const [itemData, setItemData] = useState([{}]);
    let {id} = useParams();

    useEffect(() => {
        setItemData(undefined);
        fetch(`http://localhost:5001/api/items/${id}`).then(
            response => response.json()
        ).then(
            data => {
                setItemData(data);
            }
        )
    }, []);


    return (
    <div>

        {typeof itemData === 'undefined' ? (
            <div><h3>Loading...</h3></div>
        ) : (
        <div className="row">
            <div>
                <h2 className="col s12 center">{itemData['name']}</h2>
            </div>
            <div className="col s12 m5 l5 xl4">
                <img src="https://itechstore.com.ar/wp-content/uploads/2023/02/1669130717_1726570.webp" />
            </div>
            <div>
                <h4>$ {itemData["price"]}</h4>
            </div>
            <div className="col s12 m6 l7 xl8">
                <p>The H2 chip powers a more personalized audio experience with the 2nd generation AirPods Pro with Wireless MagSafe Charging Case from Apple.
                        The chip works with a custom-built low-distortion driver, amplifier, and inward-facing mic in each ear to deliver clear highs,
                        powerful bass, and audio tuned to your ear shape as you hear it. They also usher in a new dimension of sound with immersive 
                        Spatial Audio, making you feel like you’re on stage with the band. Other aspects of the AirPods Pro have been upgraded, 
                        including ANC with double the cancellation, a transparency mode that adapts to high decibel sounds, and more. 
                        With all the reengineered features in the AirPods Pro 2nd gen, Apple takes your mobile audio experience to a new, more personal level.</p>
            </div>
        </div>
        )}
    </div>
    )
}
 
export default Item;