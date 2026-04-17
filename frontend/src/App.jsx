import { useState } from 'react';
import './App.css';

const initialNotes = [
  { id: 1, text: 'สร้างบันทึกแรกของคุณ', createdAt: 'วันนี้' },
  { id: 2, text: 'ใช้ปุ่มด้านล่างเพื่อเพิ่มบันทึกใหม่', createdAt: 'วันนี้' },
];

function App() {
  const [notes, setNotes] = useState(initialNotes);
  const [draft, setDraft] = useState('');

  const addNote = () => {
    if (!draft.trim()) return;
    setNotes((prev) => [
      { id: Date.now(), text: draft.trim(), createdAt: 'ตอนนี้' },
      ...prev,
    ]);
    setDraft('');
  };

  const removeNote = (id) => {
    setNotes((prev) => prev.filter((note) => note.id !== id));
  };

  return (
    <div className="App">
      <main className="App-card">
        <h1>Quick Note Board</h1>
        <p>เขียนไอเดียสั้น ๆ และบันทึกได้ทันที</p>

        <div className="App-input-group">
          <input
            value={draft}
            onChange={(e) => setDraft(e.target.value)}
            placeholder="พิมพ์บันทึกใหม่..."
            aria-label="New note"
          />
          <button onClick={addNote} disabled={!draft.trim()}>
            บันทึก
          </button>
        </div>

        <div className="App-notes">
          {notes.length === 0 ? (
            <p className="App-empty">ยังไม่มีบันทึก ลองเพิ่มดู</p>
          ) : (
            notes.map((note) => (
              <article key={note.id} className="App-note">
                <div>
                  <strong>{note.text}</strong>
                  <p className="App-note-meta">{note.createdAt}</p>
                </div>
                <button onClick={() => removeNote(note.id)} aria-label="Delete note">
                  ลบ
                </button>
              </article>
            ))
          )}
        </div>
      </main>
    </div>
  );
}

export default App;
