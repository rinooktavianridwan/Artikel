import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getArtikel, deleteArtikel } from "../api";

const TABS = [
  { status: "publish", label: "Published" },
  { status: "draft", label: "Drafts" },
  { status: "trash", label: "Trashed" },
];

export default function AllPosts() {
  const [posts, setPosts] = useState([]);
  const [activeTab, setActiveTab] = useState("publish");
  const [offset, setOffset] = useState(0);
  const limit = 5;

  async function loadPosts() {
    try {
      const data = await getArtikel(limit, offset, activeTab);
      setPosts(data || []);
    } catch (error) {
      console.error("Error fetching posts:", error);
    }
  }

  useEffect(() => {
    let cancelled = false;

    getArtikel(limit + 1, offset, activeTab)
      .then((data) => {
        if (!cancelled) {
          setPosts(data || []);
        }
      })
      .catch((error) => {
        console.error("Error fetching posts:", error);
      });

    return () => {
      cancelled = true;
    };
  }, [activeTab, offset, limit]);

  const hasNext = posts.length > limit;
  const visible = posts.slice(0, limit);

  async function handleTash(id) {
    try {
      await deleteArtikel(id);
      loadPosts();
    } catch (error) {
      console.error("Error deleting post:", error);
    }
  }

  return (
    <div className="p-4">
      <div className="flex justify-center mb-4">
        <h1 className="text-2xl font-bold">All Posts</h1>
      </div>

      <div className="flex gap-10 mb-4 justify-center">
        {TABS.map((tab) => (
          <button
            key={tab.status}
            onClick={() => {
              setActiveTab(tab.status);
              setOffset(0);
            }}
            className={`px-4 py-2 rounded ${
              activeTab === tab.status
                ? "bg-blue-500 text-white"
                : "bg-gray-200"
            }`}
          >
            {tab.label}
          </button>
        ))}
      </div>
      <table className="min-w-full bg-white border ">
        <thead>
          <tr>
            <th className="py-2 px-4 border">Title</th>
            <th className="py-2 px-4 border">Category</th>
            <th className="py-2 px-4 border">Actions</th>
          </tr>
        </thead>
        <tbody>
          {visible.map((post) => (
            <tr key={post.id}>
              <td className="py-2 px-4 border text-center">{post.title}</td>
              <td className="py-2 px-4 border text-center">{post.category}</td>
              <td className="py-2 px-4 border text-center">
                <Link
                  to={`/edit/${post.id}`}
                  className="text-green-500 hover:underline mr-2"
                >
                  Edit
                </Link>
                {activeTab !== "trash" && (
                  <button
                    onClick={() => handleTash(post.id)}
                    className="text-red-500 hover:underline"
                  >
                    Trash
                  </button>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {/* pagination */}
      <div className="flex justify-center mt-4 gap-2">
        <button
          onClick={() => setOffset(Math.max(0, offset - limit))}
          disabled={offset === 0}
          className="px-4 py-2 bg-gray-300 rounded disabled:opacity-50"
        >
          Previous
        </button>
        <span className="px-4 py-2">Page {Math.floor(offset / limit) + 1}</span>
        <button
          onClick={() => setOffset(offset + limit)}
          disabled={!hasNext}
          className="px-4 py-2 bg-gray-300 rounded disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  );
}
