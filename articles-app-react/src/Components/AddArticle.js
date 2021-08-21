import { useState } from "react";

function AddArticle() {
    const [title, setTitle] = useState("");
    const [desc, setDesc] = useState("");
    const [content, setContent] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        // Simple POST request with a JSON body using fetch
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ Title: title, Desc:desc, Content:content})
        };
        fetch('http://localhost:10000/article', requestOptions)
    }

    return <div>
        <form>
            <input
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder={"Write a title"}
            />
            <input
            value={desc}
            onChange={(e) => setDesc(e.target.value)}
            placeholder={"Write a description"}
            />
            <input
            value={content}
            onChange={(e) => setContent(e.target.value)}
            placeholder={"Write a content"}
            />

            <button type="submit" onClick={handleSubmit}>Validate</button>
        </form>
    </div>;
}

export default AddArticle;