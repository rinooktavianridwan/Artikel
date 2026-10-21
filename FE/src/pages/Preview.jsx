import { getArtikel } from "../api";
import { useState, useEffect } from "react";

//Preview menampilkan semua artikel yang sudah di publish dalam bentuk blog.
export default function Preview() {
  const [post, setPost] = useState(null);
  const [offset, setOffset] = useState(0);
  const limit = 5;

  useEffect(() => {
    getArtikel(limit + 1, offset, "publish")
      .then((data) => {
        setPost(data);
      })
      .catch((error) => {
        console.error("Error fetching post:", error);
      });
  }, [offset, limit]);

  if (!post) {
    return <div>Loading...</div>;
  }

  const hasNext = post.length > limit;
  const visible = post.slice(0, limit);

  return (
    <div className="max-w-4xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4 text-center">Published Artikel</h1>
      {visible.map((artikel) => (
        <div key={artikel.id} className="mb-6 border rounded-lg p-4 shadow-md">
          <h2 className="text-xl font-semibold mb-2">{artikel.title}</h2>
          <p className="text-gray-700 mb-2">{artikel.content}</p>
          <p className="text-gray-500 text-sm">Category: {artikel.category}</p>
        </div>
      ))}
      <div className="flex justify-between mt-4">
        <button
          onClick={() => setOffset(offset - limit)}
          disabled={offset === 0}
          className={`px-4 py-2 rounded ${
            offset === 0
              ? "bg-gray-300 cursor-not-allowed"
              : "bg-blue-500 text-white"
          }`}
        >
          Previous
        </button>
        <button
          onClick={() => setOffset(offset + limit)}
          disabled={!hasNext}
          className={`px-4 py-2 rounded ${
            !hasNext
              ? "bg-gray-300 cursor-not-allowed"
              : "bg-blue-500 text-white"
          }`}
        >
          Next
        </button>
      </div>
    </div>
  );
}
