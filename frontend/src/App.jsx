import React, { useState, useEffect } from "react";
import axios from "axios";

const API_URL = "/api/books";

function App() {
  const [books, setBooks] = useState([]);
  const [formData, setFormData] = useState({
    title: "",
    author: "",
    price: "",
    published_date: "",
  });
  const [editingId, setEditingId] = useState(null);

  const fetchBooks = async () => {
    try {
      const response = await axios.get(API_URL);
      setBooks(response.data);
    } catch (error) {
      console.error("Error fetching books:", error);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = {
      ...formData,
      price: parseFloat(formData.price),
      published_date: formData.published_date || null,
    };

    try {
      if (editingId) {
        await axios.put(`${API_URL}/${editingId}`, data);
      } else {
        await axios.post(API_URL, data);
      }
      setFormData({ title: "", author: "", price: "", published_date: "" });
      setEditingId(null);
      fetchBooks();
    } catch (error) {
      console.error("Error saving book:", error);
    }
  };

  const handleEdit = (book) => {
    setEditingId(book.id);
    setFormData({
      title: book.title,
      author: book.author,
      price: book.price.toString(),
      published_date: book.published_date
        ? book.published_date.split("T")[0]
        : "",
    });
  };

  const handleDelete = async (id) => {
    if (window.confirm("Delete this book?")) {
      try {
        await axios.delete(`${API_URL}/${id}`);
        fetchBooks();
      } catch (error) {
        console.error("Error deleting book:", error);
      }
    }
  };

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>Book Store</h1>

      <form onSubmit={handleSubmit} style={{ marginBottom: "20px" }}>
        <input
          placeholder="Title"
          value={formData.title}
          onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          required
        />
        <input
          placeholder="Author"
          value={formData.author}
          onChange={(e) => setFormData({ ...formData, author: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Price"
          value={formData.price}
          onChange={(e) => setFormData({ ...formData, price: e.target.value })}
          required
        />
        <input
          type="date"
          value={formData.published_date}
          onChange={(e) =>
            setFormData({ ...formData, published_date: e.target.value })
          }
        />
        <button type="submit">{editingId ? "Update" : "Add Book"}</button>
        {editingId && (
          <button onClick={() => setEditingId(null)}>Cancel</button>
        )}
      </form>

      <div style={{ display: "grid", gap: "10px" }}>
        {books.map((book) => (
          <div
            key={book.id}
            style={{
              border: "1px solid #ccc",
              padding: "10px",
              borderRadius: "8px",
            }}
          >
            <h3>{book.title}</h3>
            <p>Author: {book.author}</p>
            <p>Price: ${book.price}</p>
            <button onClick={() => handleEdit(book)}>Edit</button>
            <button onClick={() => handleDelete(book.id)}>Delete</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
