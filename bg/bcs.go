package bg

import (
	"bufio"
	"encoding/json"
	"io"
	"unicode"
)

type BCS struct {
	inner string
}

type bscTokenTypes uint32

const (
	/** End of File. */
	EOF bscTokenTypes = 0
	/** RegularExpression Id. */
	SINGLE_LINE_COMMENT bscTokenTypes = 6
	/** RegularExpression Id. */
	MULTI_LINE_COMMENT bscTokenTypes = 8
	/** RegularExpression Id. */
	IF bscTokenTypes = 10
	/** RegularExpression Id. */
	THEN bscTokenTypes = 11
	/** RegularExpression Id. */
	RESPONSE bscTokenTypes = 12
	/** RegularExpression Id. */
	END bscTokenTypes = 13
	/** RegularExpression Id. */
	NUMBER_LITERAL bscTokenTypes = 14
	/** RegularExpression Id. */
	DEC_LITERAL bscTokenTypes = 15
	/** RegularExpression Id. */
	HEX_LITERAL bscTokenTypes = 16
	/** RegularExpression Id. */
	BIN_LITERAL bscTokenTypes = 17
	/** RegularExpression Id. */
	STRING_LITERAL bscTokenTypes = 18
	/** RegularExpression Id. */
	STRING_QUOTE bscTokenTypes = 19
	/** RegularExpression Id. */
	STRING_PERCENT bscTokenTypes = 20
	/** RegularExpression Id. */
	STRING_POUND bscTokenTypes = 21
	/** RegularExpression Id. */
	STRING_TILDE bscTokenTypes = 22
	/** RegularExpression Id. */
	STRING_MULTI_TILDE bscTokenTypes = 23
	/** RegularExpression Id. */
	IDENTIFIER bscTokenTypes = 24
	/** RegularExpression Id. */
	IDENTIFIER_ESCAPED bscTokenTypes = 25
	/** RegularExpression Id. */
	LETTER bscTokenTypes = 26
	/** RegularExpression Id. */
	SPECIAL_LETTER bscTokenTypes = 27
	/** RegularExpression Id. */
	DIGIT bscTokenTypes = 28
	/** RegularExpression Id. */
	LPAREN bscTokenTypes = 29
	/** RegularExpression Id. */
	RPAREN bscTokenTypes = 30
	/** RegularExpression Id. */
	LBRACKET bscTokenTypes = 31
	/** RegularExpression Id. */
	RBRACKET bscTokenTypes = 32
	/** RegularExpression Id. */
	COMMA bscTokenTypes = 33
	/** RegularExpression Id. */
	DOT bscTokenTypes = 34
	/** RegularExpression Id. */
	BANG bscTokenTypes = 35
	/** RegularExpression Id. */
	OR bscTokenTypes = 36
	/** RegularExpression Id. */
	MINUS bscTokenTypes = 37
	/** RegularExpression Id. */
	PLUS bscTokenTypes = 38
	/** RegularExpression Id. */
	POUND bscTokenTypes = 39
)

var bcsTypeName = map[bscTokenTypes]string{
	EOF:                 "<EOF>",
	1:                   "\" \"",
	2:                   "\"\\t\"",
	3:                   "\"\\n\"",
	4:                   "\"\\r\"",
	5:                   "\"\\f\"",
	SINGLE_LINE_COMMENT: "<SINGLE_LINE_COMMENT>",
	7:                   "\"/*\"",
	MULTI_LINE_COMMENT:  "\"*/\"",
	9:                   "<token of kind 9>",
	IF:                  "\"IF\"",
	THEN:                "\"THEN\"",
	RESPONSE:            "\"RESPONSE\"",
	END:                 "\"END\"",
	"<NUMBER_LITERAL>",
	"<DEC_LITERAL>",
	"<HEX_LITERAL>",
	"<BIN_LITERAL>",
	"<STRING_LITERAL>",
	"<STRING_QUOTE>",
	"<STRING_PERCENT>",
	"<STRING_POUND>",
	"<STRING_TILDE>",
	"<STRING_MULTI_TILDE>",
	"<IDENTIFIER>",
	"<IDENTIFIER_ESCAPED>",
	"<LETTER>",
	"<SPECIAL_LETTER>",
	"<DIGIT>",
	"\"(\"",
	"\")\"",
	"\"[\"",
	"\"]\"",
	"\",\"",
	"\".\"",
	"\"!\"",
	"\"|\"",
	"\"-\"",
	"\"+\"",
	"\"#\"",
}

/*

   fn_spec_re = re.compile ('([A-Za-z0-9_]+)\s*\((.*)\)')

       self.c = None
       self.token = None
       self.ids_codes = {}


   def read_token (self, stream):
       def getc ():
           if self.c is None:
               return stream.get_char ()
           else:
               c = self.c
               self.c = None
               return c

       def nextc ():
           if self.c is None:
               self.c = stream.get_char ()
               return self.c
           else:
               return self.c

       # Skip over spaces and newlines
       c = getc ()
       while c.isspace ():
           if c == "\n":
           c = getc ()

       if c == '':
           return None

       elif c == '[':
           #res = c
           res = ''
           while True:
               c = getc ()
               if c == ']':
                   break
               res = res + c
           res2 = res.split (',')
           if len (res2) != 2:
               res2 = res.split ('.')
               if len (res2) != 4:
                   raise ValueError ("[]  len")
           res = [ int (n) for n in res2 ]
           return res

       elif c == '"':
           res = c
           while True:
               c = getc ()
               res = res + c
               if c == '"':
                   break
           return res

       elif c.isdigit () or (c == '-' and nextc ().isdigit ()):
           if c == '-':
               signum = -1
               c = getc ()
           else:
               signum = 1

           res = c
           while nextc ().isdigit ():
               res = res + getc ()
           res = signum * int (res)
           return res

       elif c.isupper ():
           res = c
           while nextc ().isalpha ():
               res = res + getc ()
           return res

       else:
           raise ValueError ("Unknown token: %s (at line %d)" %(c, self.lineno))


   def get_token (self, stream):
       if self.token is None:
           tok = self.read_token (stream)
           //p='.'
       else:
           tok = self.token
           self.token = None
           //p='x'

       //print '>>' + p+':'+repr(tok) + '<<'
       return tok

   def next_token (self, stream):
           if self.token is None:
               self.token = self.read_token (stream)
               return self.token
           else:
               return self.token

   def expect_token (self, stream, tok):
       rtok = self.get_token (stream)
       if rtok != tok:
           raise ValueError ("Expected %s, got %s (at line %d)" %(tok, str (rtok), self.lineno))
       return rtok


   def read (self, stream):
       // <SC> -> SC <SCtail>
       // <SCtail> -> <CR> <SCtail>
       // <SCtail> -> SC

       obj = []

       obj.append (self.expect_token (stream, 'SC'))
       while self.next_token (stream) == 'CR':
           obj.append (self.read_condition_response_block (stream))
       self.expect_token (stream, 'SC')

       self.script = obj
       return self


   def read_condition_response_block (self, stream):
       // <CR> -> CR <CRtail>
       // <CRtail> -> <CO> <RS> <CRtail>
       // <CRtail> -> CR

       obj = []

       obj.append (self.expect_token (stream, 'CR'))
       while self.next_token (stream) == 'CO':
           obj.append (self.read_condition (stream))
           obj.append (self.read_response_set (stream))
       self.expect_token (stream, 'CR')

       return obj


   def read_condition (self, stream):
       // <CO> -> CO <COtail>
       // <COtail> -> <TR> <COtail>
       // <COtail> -> CO

       obj = []

       obj.append (self.expect_token (stream, 'CO'))
       while self.next_token (stream) == 'TR':
           obj.append (self.read_trigger (stream))
       self.expect_token (stream, 'CO')

       return obj

   def read_trigger (self, stream):
       # pst:  id, 4*I, point, 2*S, O
       # FIXME: not correct, one of the ints is flags field
       obj = []

       obj.append (self.expect_token (stream, 'TR'))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))

       if self.type == 'pst':
           obj.append (self.get_token (stream))

       obj.append (self.read_object (stream))
       self.expect_token (stream, 'TR')

       self.add_ids_code('TRIGGER', obj[1])
       #print core.id_to_symbol ('TRIGGER', obj[1])
       return obj

   def read_response_set (self, stream):
       // <RS> -> RS <RStail>
       // <RStail> -> <RE> <RStail>
       // <RStail> -> RS

       obj = []
       obj.append (self.expect_token (stream, 'RS'))
       while self.next_token (stream) == 'RE':
           obj.append (self.read_response (stream))
       self.expect_token (stream, 'RS')

       return obj

   def read_response (self, stream):
       // <RE> -> RE int <AC> <REtail>
       // <REtail> -> <AC> <REtail>
       // <REtail> -> RE
       // FIXME: this is different from what is described in IESDP, where
       //   each response has only one action, but e.g. look at PST's 0202FD1.BCS

       obj = []
       obj.append (self.expect_token (stream, 'RE'))
       obj.append (self.get_token (stream))
       while self.next_token (stream) == 'AC':
           obj.append (self.read_action (stream))
       self.expect_token (stream, 'RE')

       return obj

   def read_action (self, stream):
       #pst:  id, 3*O, 5*I, 2*S
       obj = []
       obj.append (self.expect_token (stream, 'AC'))
       obj.append (self.get_token (stream))
       obj.append (self.read_object (stream))
       obj.append (self.read_object (stream))
       obj.append (self.read_object (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       self.expect_token (stream, 'AC')
       self.add_ids_code('ACTION', obj[1])
       return obj


   def read_object (self, stream):
       obj = []
       # PST: id, 13*I, rect, S
       obj.append (self.expect_token (stream, 'OB'))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       obj.append (self.get_token (stream))
       // print obj
       // if self.next_token (stream) == 'OB':
       //     print "short object"
       //     self.get_token (stream)
       //     return obj

       if self.type == 'pst':
           obj.append (self.get_token (stream))
           obj.append (self.get_token (stream))
           obj.append (self.get_token (stream))

       self.expect_token (stream, 'OB')

       // print core.id_to_symbol ('OBJECT', obj[10])
       return obj

   def printme (self):
       odef = {
               'pst_tr': {
                          'id': (0,),
                          'I': (1, 3),
                          'not': (2,),
                          'P': (5,),
                          'S': (6,7),
                          'O': (8,),
                          },
               'bg2_tr': { // FIXME: not tested, bogus
                          'id': (0,),
                          'I': (1, 3),
                          'not': (2,),
                          'S': (5,6),
                          'O': (7,),
                          },
               'pst_ac': {
                          'id': (0,),
                          'O': (2,3,1),
                          'I': (4, 7, 8),
                          'P': (5,), # actually, it consists of list of two integers
                          'S': (9, 10),
                          },
               }

       // FIXME: globalsetglobal, globalorglobal, ...
       split_string_fns = ('global',
                                   'globalgt',
                                   'globallt',
                                   'globalband',
                                   'globalbor',
                                   'globalmin',
                                   'globalmax',
                                   'globalshl',
                                   'globalshr',
                                   'globalxor',
                                   'bitcheck',
                                   'bitcheckexact',

                                   'globalset',
                                   'setglobal',
                                   'incrementglobal',
                                   'bitset',
                                   'bitclear')

       def resolve_object (obj):
           res = []
           if self.type == 'pst':
               if obj[1] != 0:
                   res.append (core.id_to_symbol ('EA', obj[1]))
               if obj[2] != 0:
                   res.append (core.id_to_symbol ('FACTION', obj[2]))
               if obj[3] != 0:
                   res.append (core.id_to_symbol ('TEAM', obj[3]))
               if obj[4] != 0:
                   res.append (core.id_to_symbol ('GENERAL', obj[4]))
               if obj[5] != 0:
                   res.append (core.id_to_symbol ('RACE', obj[5]))
               if obj[6] != 0:
                   res.append (core.id_to_symbol ('CLASS', obj[6]))
               if obj[7] != 0:
                   res.append (core.id_to_symbol ('SPECIFIC', obj[7]))
               if obj[8] != 0:
                   res.append (core.id_to_symbol ('GENDER', obj[8]))
               if obj[9] != 0:
                   res.append (core.id_to_symbol ('ALIGNMEN', obj[9]))
           else:
               if obj[1] != 0:
                   res.append (core.id_to_symbol ('EA', obj[1]))
               if obj[2] != 0:
                   res.append (core.id_to_symbol ('GENERAL', obj[2]))
               if obj[3] != 0:
                   res.append (core.id_to_symbol ('RACE', obj[3]))
               if obj[4] != 0:
                   res.append (core.id_to_symbol ('CLASS', obj[4]))
               if obj[5] != 0:
                   res.append (core.id_to_symbol ('SPECIFIC', obj[5]))
               if obj[6] != 0:
                   res.append (core.id_to_symbol ('GENDER', obj[6]))
               if obj[7] != 0:
                   res.append (core.id_to_symbol ('ALIGNMEN', obj[7]))

           if res:
               res = '[' + '.'.join (res) + ']'
           else:
               res = None

           res2 = None
           if self.type == 'pst':
               if obj[10] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[10])
               if obj[11] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[11]) + '(' + res2 +')'
               if obj[12] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[12]) + '(' + res2 +')'
               if obj[13] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[13]) + '(' + res2 +')'
               if obj[14] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[14]) + '(' + res2 +')'
           else:
               if obj[8] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[8])
               if obj[9] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[9]) + '(' + res2 +')'
               if obj[10] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[10]) + '(' + res2 +')'
               if obj[11] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[11]) + '(' + res2 +')'
               if obj[12] != 0:
                   res2 = core.id_to_symbol ('OBJECT', obj[12]) + '(' + res2 +')'

           # FIXME: what about area?? [-1.-1.-1.-1] (15)
           res3 = None
           if self.type == 'pst':
               if obj[16] != '""':
                   res3 = obj[16]
           else:
               if obj[13] != '""':
                   res3 = obj[13]

           if (res and res2) or (res and res3) or (res2 and res3):
               raise ValueError ("Error: More values for object: " + repr (res) + '//' + repr (res2) + '//' + repr (res3))

           if res:
               return res
           elif res2:
               return res2
           elif res3:
               return res3
           else:
               return "'OBJ"


       def resolve_args (fn_spec, obj_type, obj):
           mo = self.fn_spec_re.match (fn_spec)
           fn_name = mo.groups ()[0]
           args = mo.groups ()[1].split (",")
           arg_indices = {}
           res_args = []

           split_string_arg = fn_name.lower () in  split_string_fns

           for arg in args:
               // print 'arg', fn_spec, arg
               if arg == '':
                   break
               type, name = arg.split (':')
               name, file = name.split ('*')

               if arg_indices.has_key (type):
                   arg_indices[type] = index = arg_indices[type] + 1
               else:
                   arg_indices[type] = index = 0

               // print 'ot:', obj_type, 'type;', type, 'index:', index
               if not split_string_arg or type != 'S' or index == 0:
                   aindex = odef[self.type + '_' + obj_type][type][index] + 1
               else:
                   aindex = odef[self.type + '_' + obj_type][type][index-1] + 1


               if type == 'I':
                   v = str (obj[aindex])
                   // print 'Lookup', file, tr[index + 1]
                   if file:
                       v = core.id_to_symbol (file, obj[aindex])
                   res_args.append (v)

               elif type == 'P':
                   res_args.append ('[%d.%d]' %(obj[aindex][0], obj[aindex][1]))

               elif type == 'S':
                   if split_string_arg:
                       if index == 0:
                           res_args.append ('"' + obj[aindex][7:])
                           res_args.append (obj[aindex][0:7] + '"')
                       elif index == 1:
                           pass
                       else:
                           res_args.append (obj[aindex])
                   else:
                       res_args.append (obj[aindex])

               elif type == 'O':
                   res_args.append (resolve_object (obj[aindex]))

           return fn_name, res_args




       for cr in self.script[1:]:
           co = cr[1]
           rs = cr[2]
           print('IF')
           for tr in co[1:]:
               fn_spec = core.id_to_symbol ('TRIGGER', tr[1])
               neg = ('!', '')[not tr[3]] # FIXME: hack, use odef[]

               print fn_spec, tr
               fn_name, res_args = resolve_args (fn_spec, 'tr', tr)
               print('    ' + neg + fn_name + '(' + ','.join (res_args) + ')')
               // print '    TR', tr_name + '#0x%04x' %tr[1]
           print('THEN')
           for re in rs[1:]:
               print('    RESPONSE #%d' %re[1])
               for ac in re[2:]:
                   // print 'AC', ac[1]
                   fn_spec = core.id_to_symbol ('ACTION', ac[1])
                   fn_name, res_args = resolve_args (fn_spec, 'ac', ac)
                   // print 'res_args:', res_args
                   print('        ' + fn_name + '(' + ','.join ([ str (a) for a in res_args]) + ')')
           print('END\n')


   def add_ids_code (self, ids, id):
       ids = ids.upper ()
       if not self.ids_codes.has_key (ids):
           self.ids_codes[ids] = {}
       if not self.ids_codes[ids].has_key (id):
           self.ids_codes[ids][id] = 1
       else:
           self.ids_codes[ids][id] += 1

   def uses_ids_code (self, ids, code):
       try:
           return self.ids_codes[ids.upper ()][code]
       except:
           return False
*/

/*
res = ”
while True:

	c = getc ()
	if c == ']':
	    break
	res = res + c

res2 = res.split (',')
if len (res2) != 2:

	res2 = res.split ('.')
	if len (res2) != 4:
	raise ValueError ("[]  len")

return [ int (n) for n in res2 ]
*/
func parseBrackets(reader *bufio.Reader) (*[]string, error) {
	out := []string{}
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}
		if r == ']' {
			break
		}
		out = append(out, string(r))
	}
	return &out, nil
}

func OpenBcs(reader io.ReadSeeker) (*BCS, error) {
	bcs := BCS{""}
	br := bufio.NewReader(reader)
	for {
		r, _, err := br.ReadRune()
		if err != nil {
			break
		}
		if unicode.IsControl(r) || unicode.IsSpace(r) {
			bcs.inner += string(r)
		}
		//    elif c == '[':
		//    #res = c
		//    res = ''
		//    while True:
		//        c = getc()
		//        if c == ']':
		//            break
		//        res = res + c
		//    res2 = res.split (',')
		//    if len (res2) != 2:
		//        res2 = res.split ('.')
		//        if len (res2) != 4:
		//            raise ValueError ("[]  len")
		//    res = [ int (n) for n in res2 ]
		//    return res
		if r == '[' {
			buff, err := parseBrackets(br)
			if err != nil {
				return nil, err
			}
			for _, part := range *buff {
				bcs.inner += part
			}

		}
	}
	return &bcs, nil
}

func (bcs *BCS) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(bcs, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
