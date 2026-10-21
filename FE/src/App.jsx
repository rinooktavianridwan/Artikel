import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import AllPosts from "./pages/AllPost";
import AddNew from "./pages/AddNew";
import EditArtikel from "./pages/EditArtikel";
import Preview from "./pages/Preview";

function App() {
  return (
    <Router>
      <nav className="flex gap-4 p-4 bg-gray-800 text-white">
        <Link to="/">All Posts</Link>
        <Link to="/add">Add New</Link>
        <Link to="/view">View</Link>
      </nav>
      <Routes>
        <Route path="/" element={<AllPosts />} />
        <Route path="/add" element={<AddNew />} />
        <Route path="/edit/:id" element={<EditArtikel />} />
        <Route path="/view" element={<Preview />} />
      </Routes>
    </Router>
  );
}

export default App;
