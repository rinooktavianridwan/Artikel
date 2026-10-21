import { useState } from "react";
import { createArtikel } from "../api";

export default function AddNew() {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("");

  const save = (status) => {
    const newPost = {
      title: title,
      content: content,
      category: category,
      status: status,
    };

    createArtikel(newPost)
      .then((response) => {
        console.log("Post created:", response);
        setTitle("");
        setContent("");
        setCategory("");
      })
      .catch((error) => {
        console.error("Error creating post:", error);
      });
  };

  return (
    <div className="max-w-2xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Add New Post</h1>
      <form>
        <div className="mb-4">
          <label
            htmlFor="title"
            className="block text-sm font-medium text-gray-700"
          >
            Title
          </label>
          <input
            type="text"
            id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="content"
            className="block text-sm font-medium text-gray-700"
          >
            Content
          </label>
          <textarea
            id="content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="category"
            className="block text-sm font-medium text-gray-700"
          >
            Category
          </label>
          <input
            type="text"
            id="category"
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <button
          type="button"
          onClick={() => save("publish")}
          className="cursor-pointer rounded bg-green-600 px-4 py-2 text-white hover:bg-green-700"
        >
          Publish
        </button>
        <button
          type="button"
          onClick={() => save("draft")}
          className="ml-2 cursor-pointer rounded bg-gray-500 px-4 py-2 text-white hover:bg-gray-600"
        >
          Draft
        </button>
      </form>
    </div>
  );
}
