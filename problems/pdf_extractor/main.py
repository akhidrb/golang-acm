from pdfquery import PDFQuery


if __name__ == '__main__':
    pdf = PDFQuery('gtg.pdf')
    pdf.load()

    # Use CSS-like selectors to locate the elements
    text_elements = pdf.pq('LTTextLineHorizontal')

    # Extract the text from the elements
    text = [t.text for t in text_elements]

    print(text)
