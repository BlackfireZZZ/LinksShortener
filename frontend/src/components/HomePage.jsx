import React, { useState } from 'react';
import axios from 'axios';
import config from "../config";

const HomePage = () => {
    const [link, setLink] = useState('');
    const [shortLink, setShortLink] = useState('');
    const [error, setError] = useState('');

    const validateLink = (url) => {
        const regex = /^(https?:\/\/)?([\w\-]+\.)+[\w\-]+(\/[\w\-]*)*\/?$/;
        return regex.test(url);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!validateLink(link)) {
            setError('Please enter a valid URL.');
            return;
        }

        setError('');

        try {
            const response = await axios.post(config.Host_url, {
                full_link: link,
            });
            setShortLink(response.data.short_link);
        } catch (err) {
            setError('Error shortening the link. Please try again.');
        }
    };

    return (
        <div>
            {shortLink ? (
                <div className="mb-3">
                    <label className="form-label">Your short link:</label>
                    <div className="alert alert-success" role="alert">
                        {shortLink}
                    </div>
                </div>
            ) : (
                <form onSubmit={handleSubmit}>
                    <div className="mb-3">
                        <label htmlFor="formGroupExampleInput" className="form-label">Enter a full link and get a short link</label>
                        <input
                            type="text"
                            className="form-control"
                            id="formGroupExampleInput"
                            placeholder="Put your link here"
                            value={link}
                            onChange={(e) => setLink(e.target.value)}
                        />
                        {error && <div className="text-danger mt-2">{error}</div>}
                    </div>
                    <button type="submit" className="btn btn-primary">Submit</button>
                </form>
            )}
        </div>
    );
};

export default HomePage;
