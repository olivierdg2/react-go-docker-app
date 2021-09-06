import { useState } from "react";
import { useHistory } from "react-router-dom";

function AddArticle() {
    const [title, setTitle] = useState("");
    const [desc, setDesc] = useState("");
    const [content, setContent] = useState("");
    const history = useHistory();
    const handleSubmit = (e) => {
        e.preventDefault();
        // Simple POST request with a JSON body using fetch
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify({ Title: title, Desc:desc, Content:content}),
            mode: "no-cors"
        };
        fetch('http://10.1.0.28:10000/article', requestOptions)
        .then(function(response) {
            history.push("/articles");
        });
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