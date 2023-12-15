import React, { useState } from 'react';
import "../style/App1.css"
import axios from "axios";

function App1() {
    const [text, setText] = useState('');
    const [score, setScore] = useState(null);
    const [history, setHistory] = useState([]);
    const [selectedPage, setSelectedPage] = useState(1); // 選択されたページ
    const [itemsPerPage] = useState(10); // 1ページに表示するアイテム数

    const analyzeSentiment = () => {
        // テキスト情報をJSONに変換

        const jsonText = JSON.stringify({ text: text });
        console.log(jsonText)
        // axiosを使用してエンドポイントにPOSTリクエストを送信
        axios.post('http://localhost:8080/sentiment_text', jsonText)
        .then(response => {
            // リクエストが成功した場合の処理
            console.log('API Response:', response.data);
            const sentimentScore = response.data.sentiment_text.sentiment_score;
            const id = response.data.sentiment_text.id;
            setScore(sentimentScore);
            setHistory([...history, { id: id, text: text, score: sentimentScore }]);
        })
        .catch(error => {
            // エラーが発生した場合の処理
            console.error('API Error:', error);
        });
    };

    const deleteItem = (index) => {
        // 指定されたインデックスのアイテムを削除
        const newHistory = [...history];
        newHistory.splice(index, 1);
        setHistory(newHistory);
    };

    // 現在のページのアイテムの範囲を計算
    const indexOfLastItem = selectedPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = history.slice(indexOfFirstItem, indexOfLastItem);

    // ページ数を表示する関数
    const renderPageOptions = () => {
        const pageOptions = [];
        for (let i = 1; i <= Math.ceil(history.length / itemsPerPage); i++) {
            pageOptions.push(
                <option key={i} value={i}>
                    Page {i}
                </option>
            );
        }
        if (currentItems.length < itemsPerPage) {
            const remainingItems = history.slice(
                selectedPage * itemsPerPage,
                (selectedPage + 1) * itemsPerPage
            );
            currentItems.push(...remainingItems);
        }
        return pageOptions;
    };


    return (
        <div className="app1">
            <div className="input-form">
        <textarea
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="テキストを入力..."
        />
                <button onClick={analyzeSentiment}>Analyze</button>
                <div className="score">
                    {score !== null && <p>Sentiment Score: {score.toFixed(2)}</p>}
                </div>
            </div>
            <div className="display-area">
                <h2>History</h2>
                <div className="pagination">
                    <select
                        value={selectedPage}
                        onChange={(e) => setSelectedPage(parseInt(e.target.value))}
                    >
                        {renderPageOptions()}
                    </select>
                </div>
                <ul className="history-list">
                    {currentItems.map((item, index) => (
                        <li key={index}>
                            Id: {item.id}, Text:{item.text}, Score: {item.score.toFixed(2)}
                            <button onClick={() => deleteItem(index)}>Delete</button>
                        </li>
                    ))}
                </ul>

            </div>
            {/* JSONデータを表示 */}
            {/*{submitText && (*/}
            {/*    <div className="json-output">*/}
            {/*        <h2>JSON Data to Submit</h2>*/}
            {/*        <pre>{submitText}</pre>*/}
            {/*    </div>*/}
            {/*)}*/}
        </div>
    );
}

export default App1;