import React from 'react';
import axios from "axios";

const App3 = () => {
    const handleSubmit = event => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        axios.post('http://localhost:8080/sentiment_text', {
            // axiosならJSONデータをリテラルで書ける
            sentiment_text: data.get("name")
        })
        .then(response => {
            // リクエストが成功した場合の処理
            console.log('API Response:', response.data);
        })
        .catch(error => {
            // エラーが発生した場合の処理
            console.error('API Error:', error);
        });
    };

    return (
        <form onSubmit={event => handleSubmit(event)}>
            <label htmlFor="name">Name: </label>
            <br/>
            <input type="text" id="name" name="name"/>
            <br/>
            <input type="submit" defaultValue={"Submit"}/>
        </form>
    );
}

export default App3;