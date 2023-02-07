function darkPlusHighlightStyle() {
  const config = {
    name: 'darkPlus',
    dark: true,
    background: '#1E1E1E',
    foreground: '#BBBBBB',
    selection: '#264F78',
    cursor: '#BBBBBB',
    dropdownBackground: '#252526',
    dropdownBorder: '#454545',
    activeLine: '#264F78',
    matchingBracket: '#0064001a',
    keyword: '#569cd6',
    storage: '#569cd6',
    variable: '#9CDCFE',
    parameter: '#9CDCFE',
    function: '#DCDCAA',
    string: '#ce9178',
    constant: '#569cd6',
    type: '#4EC9B0',
    class: '#4EC9B0',
    number: '#b5cea8',
    comment: '#6a9955',
    heading: '#000080',
    invalid: '#f44747',
    regexp: '#646695',
  };
  return HighlightStyle.define([ 
    {tag: t.keyword, color: config.keyword},
    {tag: [t.name, t.deleted, t.character, t.macroName], color: config.variable},
    {tag: [t.propertyName], color: config.function},
    {tag: [t.processingInstruction, t.string, t.inserted, t.special(t.string)], color: config.string},
    {tag: [t.function(t.variableName), t.labelName], color: config.function},
    {tag: [t.color, t.constant(t.name), t.standard(t.name)], color: config.constant},
    {tag: [t.definition(t.name), t.separator], color: config.variable},
    {tag: [t.className], color: config.class},
    {tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace], color: config.number},
    {tag: [t.typeName], color: config.type, fontStyle: config.type},
    {tag: [t.operator, t.operatorKeyword], color: config.keyword},
    {tag: [t.url, t.escape, t.regexp, t.link], color: config.regexp},
    {tag: [t.meta, t.comment], color: config.comment},
    {tag: t.strong, fontWeight: 'bold'},
    {tag: t.emphasis, fontStyle: 'italic'},
    {tag: t.link, textDecoration: 'underline'},
    {tag: t.heading, fontWeight: 'bold', color: config.heading},
    {tag: [t.atom, t.bool, t.special(t.variableName)], color: config.variable},
    {tag: t.invalid, color: config.invalid},
    {tag: t.strikethrough, textDecoration: 'line-through'},
  ]);
}
